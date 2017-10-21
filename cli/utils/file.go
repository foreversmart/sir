package utils

import (
	"os"
	"os/exec"
	"path/filepath"
)

func ExecFileAbsPath(path string) (abspath string, err error) {
	abspath, err = exec.LookPath(path)
	if err == nil {
		return
	}

	abspath, err = filepath.Abs(path)
	if err != nil {
		return
	}

	_, err = os.Stat(path)
	if err != nil {
		return
	}

	return
}
