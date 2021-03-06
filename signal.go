package raft

import (
	"os"
	"os/signal"
	"syscall"
)

// terminalSignalCh returns a channel that waits for signals which usually indicates
// the terminal of a process.
func terminalSignalCh() <-chan os.Signal {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	return ch
}
