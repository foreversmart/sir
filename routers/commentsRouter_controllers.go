package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["sir/controllers:TaskController"] = append(beego.GlobalControllerRouter["sir/controllers:TaskController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/task`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["sir/controllers:TaskController"] = append(beego.GlobalControllerRouter["sir/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Show",
			Router: `/task/:name`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["sir/controllers:TaskController"] = append(beego.GlobalControllerRouter["sir/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/task/:name`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["sir/controllers:TaskController"] = append(beego.GlobalControllerRouter["sir/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Remove",
			Router: `/task/:name`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["sir/controllers:TaskController"] = append(beego.GlobalControllerRouter["sir/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Log",
			Router: `/task/:name/log`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["sir/controllers:TaskController"] = append(beego.GlobalControllerRouter["sir/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Rename",
			Router: `/task/:name/rename`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["sir/controllers:TaskController"] = append(beego.GlobalControllerRouter["sir/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Restart",
			Router: `/task/:name/restart`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["sir/controllers:TaskController"] = append(beego.GlobalControllerRouter["sir/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Send",
			Router: `/task/:name/send`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["sir/controllers:TaskController"] = append(beego.GlobalControllerRouter["sir/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Start",
			Router: `/task/:name/start`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["sir/controllers:TaskController"] = append(beego.GlobalControllerRouter["sir/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Statistics",
			Router: `/task/:name/statistics`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["sir/controllers:TaskController"] = append(beego.GlobalControllerRouter["sir/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Stop",
			Router: `/task/:name/stop`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["sir/controllers:TaskController"] = append(beego.GlobalControllerRouter["sir/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/task/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["sir/controllers:TaskController"] = append(beego.GlobalControllerRouter["sir/controllers:TaskController"],
		beego.ControllerComments{
			Method: "AllLog",
			Router: `/task/log`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
