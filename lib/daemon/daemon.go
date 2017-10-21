package daemon

import (
	"fmt"
	"os"
	"syscall"
)

func Daemon(stdLogPath string, errLogPath string) (err error) {
	// check it's ready a daemon
	if syscall.Getppid() == 1 {
		// set unmask 0
		syscall.Umask(0)
		os.Chdir("/")
		return nil
	}

	// set output flow
	files := make([]*os.File, 3)
	stdLogFile, err := os.OpenFile(stdLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	errLogFile, err := os.OpenFile(errLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	files[0], files[1], files[2] = os.Stdin, stdLogFile, errLogFile

	// work dir
	dir, _ := os.Getwd()
	// set set id
	attr := syscall.SysProcAttr{Setsid: true}
	procAttrs := os.ProcAttr{Dir: dir, Env: os.Environ(), Files: files, Sys: &attr}

	// start child process
	proc, err := os.StartProcess(os.Args[0], os.Args, &procAttrs)
	if err != nil {
		return fmt.Errorf("can't create process %s: %s", os.Args[0], err)
	}
	proc.Release()
	os.Exit(0)

	return nil
}
