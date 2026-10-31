//go:build !windows

package store

import "syscall"

func setUmask(mask int) int {
	return syscall.Umask(mask)
}

func restoreUmask(mask int) {
	syscall.Umask(mask)
}

