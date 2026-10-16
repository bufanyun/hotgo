//go:build unix

package queue

import (
	"os"
	"syscall"
)

// lockFile 对整个文件加排他锁
func lockFile(f *os.File) error {
	return syscall.Flock(int(f.Fd()), syscall.LOCK_EX)
}

// unlockFile 释放锁
func unlockFile(f *os.File) error {
	return syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
}
