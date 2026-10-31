//go:build windows

package cmd

import (
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func detachedSysProcAttr() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{
		CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP,
	}
}

func execProcess(binary string, args []string, env []string) error {
	c := exec.Command(binary, args[1:]...)
	c.Env = env
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			os.Exit(exitErr.ExitCode())
		}
		return err
	}
	os.Exit(0)
	return nil
}

func terminateProcess(proc *os.Process) error {
	return proc.Kill()
}

func notifyContainerSignals(sigs chan os.Signal) {
	signal.Notify(sigs, os.Interrupt)
}

