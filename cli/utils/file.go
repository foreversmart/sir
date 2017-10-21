package utils

import (
	"os"
	"os/exec"
	"path/filepath"
)

// 获取可执行文件的绝对路径
func ExecFileAbsPath(path string) (abspath string, err error) {
	abspath, err = exec.LookPath(path)
	if err == nil && abspath[0] != '.' {
		return
	}

	abspath, err = filepath.Abs(path)
	if err != nil {
		return
	}

	_, err = os.Stat(abspath)
	if err != nil {
		return
	}

	return
}
