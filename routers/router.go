package routers

import (
	"sir/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.TaskController{})
	beego.SetStaticPath("/view", "public")
}
