package controllers

import (
	"io/ioutil"
	"os"
	"sir/lib/errors"
	"sir/models"

	"log"

	"github.com/BurntSushi/toml"
)

var (
	taskConfigPath = "~/.sir/taskconfig"
)

func CreateConfig(params *models.TaskConfig) (err error) {
	if !params.IsValid() {
		return errors.InvalidTaskConfig
	}

	file, err := os.OpenFile(taskConfigPath+"/"+params.Name+".toml", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Printf("os.OpenFile(%s, os.O_RDWR, 0666): %v", taskConfigPath)

		return
	}

	err = toml.NewEncoder(file).Encode(params)
	if err != nil {
		log.Printf("toml.NewEncoder(%s).Encode(%#v): %v", file, params)

		return
	}

	return
}

func GetConfig(appName string) (config *models.TaskConfig, err error) {
	config = &models.TaskConfig{}

	files, err := ioutil.ReadDir(taskConfigPath)
	if err != nil {
		log.Printf("ioutil.ReadDir(%s): %v", taskConfigPath, err)

		return
	}

	for _, file := range files {
		var conf *models.TaskConfig

		_, err = toml.DecodeFile(taskConfigPath+"/"+file.Name(), conf)
		if err != nil {
			log.Printf("toml.DecodeFile(%s, conf): %v", taskConfigPath+"/"+file.Name(), err)
		}

		if conf.Name == appName {
			config = conf
			break
		}
	}

	return
}
