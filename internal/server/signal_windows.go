//go:build windows

package server

import (
	"os"
	"os/signal"
)

func notifyShutdown(stop chan os.Signal) {
	signal.Notify(stop, os.Interrupt)
}

