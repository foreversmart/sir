package main

import (
	"github.com/astaxie/beego"
	"sir/lib/daemon"
	_ "sir/routers"
)

func main() {
	err := daemon.Daemon("/Users/hong/projects/src/sir/std.log", "/Users/hong/projects/src/sir/err.log")
	beego.Error(err)

	beego.Run()
}
