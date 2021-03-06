package raft

import (
	"context"
	"encoding/base64"
	"errors"
	"io"
	"log"
	"net"
	"net/rpc"
	"sync"
	"sync/atomic"

	"github.com/sumimakito/raft/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

type grpcTransService struct {
	rpcCh chan *RPC
	pb.UnimplementedTransportServer
}

func (s *grpcTransService) AppendEntries(ctx context.Context, request *pb.AppendEntriesRequest) (*pb.AppendEntriesResponse, error) {
	r := NewRPC(ctx, request)
	s.rpcCh <- r
	response, err := r.Response()
	if err != nil {
		return nil, err
	}
	return response.(*pb.AppendEntriesResponse), nil
}

func (s *grpcTransService) RequestVote(ctx context.Context, request *pb.RequestVoteRequest) (*pb.RequestVoteResponse, error) {
	r := NewRPC(ctx, request)
	s.rpcCh <- r
	response, err := r.Response()
	if err != nil {
		return nil, err
	}
	return response.(*pb.RequestVoteResponse), nil
}

func (s *grpcTransService) InstallSnapshot(stream pb.Transport_InstallSnapshotServer) error {
	streamMetadata, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return errors.New("invalid metadata")
	}
	var requestMetaBase64 string
	if values := streamMetadata.Get("requestMeta"); len(values) < 1 {
		return errors.New("invalid metadata")
	} else {
		requestMetaBase64 = values[0]
	}
	requestMetaBytes, err := base64.StdEncoding.DecodeString(requestMetaBase64)
	if err != nil {
		return err
	}
	var requestMeta pb.InstallSnapshotRequestMeta
	if err := proto.Unmarshal(requestMetaBytes, &requestMeta); err != nil {
		return err
	}

	pr, pw := io.Pipe()
	writer := NewBufferedWriteCloser(pw)

	request := &InstallSnapshotRequest{
		Metadata: &requestMeta,
		Reader:   NewBufferedReadCloser(pr),
	}

	r := NewRPC(stream.Context(), request)
	s.rpcCh <- r

	go func() {
		defer writer.Close()
		for {
			requestData, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				r.Respond(nil, err)
				return
			}
			if _, err := writer.Write(requestData.Data); err != nil {
				r.Respond(nil, err)
				return
			}
		}
		writer.Flush()
	}()

	response, err := r.Response()
	if err != nil {
		return err
	}
	return stream.SendAndClose(response.(*pb.InstallSnapshotResponse))
}

func (s *grpcTransService) ApplyLog(ctx context.Context, request *pb.ApplyLogRequest) (*pb.ApplyLogResponse, error) {
	r := NewRPC(ctx, request)
	s.rpcCh <- r
	response, err := r.Response()
	if err != nil {
		return nil, err
	}
	return response.(*pb.ApplyLogResponse), nil
}

type grpcTransClient struct {
	conn   *grpc.ClientConn
	client pb.TransportClient
}

type GRPCTransport struct {
	service *grpcTransService
	server  *grpc.Server

	listener net.Listener

	serveFlag uint32

	clients   map[string]*grpcTransClient
	clientsMu sync.RWMutex // protects clients
}

func NewGRPCTransport(listenAddr string) (*GRPCTransport, error) {
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return nil, err
	}
	return &GRPCTransport{
		service:  &grpcTransService{rpcCh: make(chan *RPC, 16)},
		listener: listener,
		clients:  map[string]*grpcTransClient{},
	}, nil
}

func (t *GRPCTransport) connectLocked(peer *pb.Peer) error {
	if _, ok := t.clients[peer.Id]; ok {
		return nil
	}
	conn, err := grpc.Dial(peer.Endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	log.Println("peer connected", "target", conn.Target())
	t.clients[peer.Id] = &grpcTransClient{conn: conn, client: pb.NewTransportClient(conn)}
	return nil
}

func (t *GRPCTransport) disconnectLocked(peer *pb.Peer) {
	if client, ok := t.clients[peer.Id]; ok {
		delete(t.clients, peer.Id)
		client.conn.Close()
	}
}

func (t *GRPCTransport) tryClient(peer *pb.Peer, fn func(c *grpcTransClient) error) error {
	retryState := -1
	var lastErr error
	var client *grpcTransClient
	var ok bool
retryClient:
	if retryState > 0 {
		return lastErr
	}
	retryState++
	t.clientsMu.RLock()
	client, ok = t.clients[peer.Id]
	t.clientsMu.RUnlock()
	// Check if the client is unset
	if !ok {
		t.clientsMu.Lock()
		// Check again to ensure the client is unset
		client, ok = t.clients[peer.Id]
		if ok {
			// Client is set
			t.clientsMu.Unlock()
			goto tryCall
		}
		// Client is unset
		// Try to connect it
		if err := t.connectLocked(peer); err != nil {
			t.clientsMu.Unlock()
			return err
		}
		t.clientsMu.Unlock()
		lastErr = errors.New("client not connected")
		goto retryClient
	}
tryCall:
	if err := fn(client); err != nil {
		if err == rpc.ErrShutdown {
			// Disconnect current client
			t.clientsMu.Lock()
			t.disconnectLocked(peer)
			// And try to connect it again
			if err := t.connectLocked(peer); err != nil {
				t.clientsMu.Unlock()
				return err
			}
			t.clientsMu.Unlock()
			lastErr = err
			goto retryClient
		}
		return err
	}
	return nil
}

func (t *GRPCTransport) Endpoint() string {
	return t.listener.Addr().String()
}

func (t *GRPCTransport) AppendEntries(
	ctx context.Context, peer *pb.Peer, request *pb.AppendEntriesRequest,
) (*pb.AppendEntriesResponse, error) {
	var response *pb.AppendEntriesResponse
	if err := t.tryClient(peer, func(c *grpcTransClient) error {
		r, err := c.client.AppendEntries(ctx, request)
		if err != nil {
			return err
		}
		response = r
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}

func (t *GRPCTransport) RequestVote(
	ctx context.Context, peer *pb.Peer, request *pb.RequestVoteRequest,
) (*pb.RequestVoteResponse, error) {
	var response *pb.RequestVoteResponse
	if err := t.tryClient(peer, func(c *grpcTransClient) error {
		r, err := c.client.RequestVote(ctx, request)
		if err != nil {
			return err
		}
		response = r
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}

func (t *GRPCTransport) InstallSnapshot(
	ctx context.Context, peer *pb.Peer, requestMeta *pb.InstallSnapshotRequestMeta, reader io.Reader,
) (*pb.InstallSnapshotResponse, error) {
	var response *pb.InstallSnapshotResponse
	if err := t.tryClient(peer, func(c *grpcTransClient) error {
		reqestMetaByets, err := proto.Marshal(requestMeta)
		if err != nil {
			return err
		}
		ctx := metadata.AppendToOutgoingContext(ctx, "requestMeta", base64.StdEncoding.EncodeToString(reqestMetaByets))
		client, err := c.client.InstallSnapshot(ctx)
		if err != nil {
			return err
		}
		chunk := make([]byte, 4096)
		for {
			n, err := reader.Read(chunk)
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}
			if err := client.Send(&pb.InstallSnapshotRequestData{Data: chunk[:n]}); err != nil {
				return err
			}
		}
		r, err := client.CloseAndRecv()
		if err != nil {
			return err
		}
		response = r
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}

func (t *GRPCTransport) ApplyLog(
	ctx context.Context, peer *pb.Peer, request *pb.ApplyLogRequest,
) (*pb.ApplyLogResponse, error) {
	var response *pb.ApplyLogResponse
	if err := t.tryClient(peer, func(c *grpcTransClient) error {
		r, err := c.client.ApplyLog(ctx, request)
		if err != nil {
			return err
		}
		response = r
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}

func (t *GRPCTransport) RPC() <-chan *RPC {
	return t.service.rpcCh
}

func (t *GRPCTransport) Serve() error {
	if !atomic.CompareAndSwapUint32(&t.serveFlag, 0, 1) {
		panic("Serve() should be only called once")
	}
	log.Println("transport started", "addr", t.listener.Addr())
	t.server = grpc.NewServer()
	pb.RegisterTransportServer(t.server, t.service)
	return t.server.Serve(t.listener)
}

func (t *GRPCTransport) Connect(peer *pb.Peer) error {
	t.clientsMu.RLock()
	if _, ok := t.clients[peer.Id]; ok {
		return nil
	}
	t.clientsMu.RUnlock()
	t.clientsMu.Lock()
	defer t.clientsMu.Unlock()
	return t.connectLocked(peer)
}

func (t *GRPCTransport) Disconnect(peer *pb.Peer) {
	t.clientsMu.Lock()
	defer t.clientsMu.Unlock()
	if client, ok := t.clients[peer.Id]; ok {
		delete(t.clients, peer.Id)
		client.conn.Close()
	}
}

func (t *GRPCTransport) DisconnectAll() {
	t.clientsMu.Lock()
	defer t.clientsMu.Unlock()
	for _, client := range t.clients {
		client.conn.Close()
	}
	t.clients = map[string]*grpcTransClient{}
}

func (t *GRPCTransport) Close() error {
	t.DisconnectAll()
	t.server.GracefulStop()
	return nil
}
