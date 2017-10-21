package main

import (
	"os"
	"sir/lib/config"
	_ "sir/routers"

	"github.com/astaxie/beego"
)

func main() {
	appHome := config.AppHome
	os.MkdirAll(appHome, 0700)
	// TODO handle error
	errLog := appHome + "/" + beego.AppConfig.String("err_log")
	stdLog := appHome + "/" + beego.AppConfig.String("std_log")

	if !Exists(errLog) {
		os.Create(errLog)
	}
	if !Exists(stdLog) {
		os.Create(stdLog)
	}

	// err := daemon.Daemon(stdLog, errLog)
	// beego.Error(err)

	beego.Run()
}

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
