package routers

import (
	"github.com/foreversmart/sir/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.TaskController{})
	beego.SetStaticPath("/public", "public")
}
