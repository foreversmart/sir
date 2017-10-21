package config

import (
	"os"
	"os/user"
	"runtime"
	"sir/lib/errors"
	"sir/models"
	"time"

	"log"

	"github.com/BurntSushi/toml"
)

var (
	AppHome        = UserHomeDir() + "/.sir"
	taskConfigPath = UserHomeDir() + "/.sir/configs"
	defaultLogPath = UserHomeDir() + "/.sir/logs"
)

func CreateTaskConfig(params *models.TaskConfig) (err error) {

	if !params.IsValid() {
		return errors.InvalidTaskConfig
	}

	params.CTime = time.Now()
	if params.LogConfigs.ErrLogPath == "" {
		params.LogConfigs.ErrLogPath = defaultLogPath
	}
	if params.LogConfigs.StdLogPath == "" {
		params.LogConfigs.StdLogPath = defaultLogPath
	}
	if params.LogConfigs.RotateType == "" {
		params.LogConfigs.RotateType = "day"
	}

	if params.User == "" {
		user, _ := user.Current()
		params.User = user.Username
	}
	if params.Group == "" {
		params.Group = params.User
	}

	err = os.MkdirAll(taskConfigPath, 0700)
	if err != nil {
		log.Printf("os.MkdirAll(%s, os.O_RDWR, 0666): %v", taskConfigPath, err)

		return
	}

	file, err := os.OpenFile(taskConfigPath+"/"+params.Name+".toml", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Printf("os.OpenFile(%s, os.O_RDWR, 0666): %v", taskConfigPath+"/"+params.Name+".toml", err)

		return
	}

	err = toml.NewEncoder(file).Encode(params)
	if err != nil {
		log.Printf("toml.NewEncoder(%s).Encode(%#v): %v", file, params)

		return
	}

	return
}

func GetTaskConfig(taskName string) (config *models.TaskConfig, err error) {
	config = &models.TaskConfig{}

	_, err = toml.DecodeFile(taskConfigPath+"/"+taskName+".toml", config)
	if err != nil {
		log.Printf("toml.DecodeFile(%s, conf): %v", taskConfigPath+"/"+taskName+".toml", err)
	}

	return
}

func DeleteTaskConfig(taskName string) error {
	return os.Remove(taskConfigPath + "/" + taskName + ".toml")
}

func GetTaskConfigFilePath(taskName string) string {
	return taskConfigPath + "/" + taskName + ".toml"
}

func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}
