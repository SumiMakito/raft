// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: transport.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_transport_proto protoreflect.FileDescriptor

var file_transport_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x09, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0x94, 0x02, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x44,
	0x0a, 0x0d, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12,
	0x18, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x41,
	0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56,
	0x6f, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0f, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x53,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x35, 0x0a, 0x08, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x4c, 0x6f, 0x67, 0x12, 0x13, 0x2e, 0x70,
	0x62, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x4c, 0x6f, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x75, 0x6d, 0x69, 0x6d, 0x61, 0x6b, 0x69, 0x74, 0x6f,
	0x2f, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_transport_proto_goTypes = []interface{}{
	(*AppendEntriesRequest)(nil),    // 0: pb.AppendEntriesRequest
	(*RequestVoteRequest)(nil),      // 1: pb.RequestVoteRequest
	(*InstallSnapshotRequest)(nil),  // 2: pb.InstallSnapshotRequest
	(*ApplyLogRequest)(nil),         // 3: pb.ApplyLogRequest
	(*AppendEntriesResponse)(nil),   // 4: pb.AppendEntriesResponse
	(*RequestVoteResponse)(nil),     // 5: pb.RequestVoteResponse
	(*InstallSnapshotResponse)(nil), // 6: pb.InstallSnapshotResponse
	(*ApplyLogResponse)(nil),        // 7: pb.ApplyLogResponse
}
var file_transport_proto_depIdxs = []int32{
	0, // 0: pb.Transport.AppendEntries:input_type -> pb.AppendEntriesRequest
	1, // 1: pb.Transport.RequestVote:input_type -> pb.RequestVoteRequest
	2, // 2: pb.Transport.InstallSnapshot:input_type -> pb.InstallSnapshotRequest
	3, // 3: pb.Transport.ApplyLog:input_type -> pb.ApplyLogRequest
	4, // 4: pb.Transport.AppendEntries:output_type -> pb.AppendEntriesResponse
	5, // 5: pb.Transport.RequestVote:output_type -> pb.RequestVoteResponse
	6, // 6: pb.Transport.InstallSnapshot:output_type -> pb.InstallSnapshotResponse
	7, // 7: pb.Transport.ApplyLog:output_type -> pb.ApplyLogResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_transport_proto_init() }
func file_transport_proto_init() {
	if File_transport_proto != nil {
		return
	}
	file_rpc_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_transport_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transport_proto_goTypes,
		DependencyIndexes: file_transport_proto_depIdxs,
	}.Build()
	File_transport_proto = out.File
	file_transport_proto_rawDesc = nil
	file_transport_proto_goTypes = nil
	file_transport_proto_depIdxs = nil
}
