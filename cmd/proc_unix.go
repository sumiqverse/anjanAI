//go:build !windows

package cmd

import (
	"os"
	"os/signal"
	"syscall"
)

func detachedSysProcAttr() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{Setsid: true}
}

func execProcess(binary string, args []string, env []string) error {
	return syscall.Exec(binary, args, env)
}

func terminateProcess(proc *os.Process) error {
	return proc.Signal(syscall.SIGTERM)
}

func notifyContainerSignals(sigs chan os.Signal) {
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
}

