//go:build !windows

package server

import (
	"os"
	"os/signal"
	"syscall"
)

func notifyShutdown(stop chan os.Signal) {
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
}

