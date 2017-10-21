package controllers

import (
	"os"
	"sir/lib/errors"
	"sir/models"

	"log"

	"github.com/BurntSushi/toml"
)

var (
	taskConfigPath = "~/.sir/taskconfig"
)

func CreateTaskConfig(params *models.TaskConfig) (err error) {
	if !params.IsValid() {
		return errors.InvalidTaskConfig
	}

	file, err := os.OpenFile(taskConfigPath+"/"+params.Name+".toml", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Printf("os.OpenFile(%s, os.O_RDWR, 0666): %v", taskConfigPath+"/"+params.Name+".toml")

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
