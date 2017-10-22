package routers

import (
	"sir/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.SetStaticPath("/public", "public")
	beego.Include(&controllers.TaskController{})
}
