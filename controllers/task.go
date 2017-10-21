package controllers

import "github.com/astaxie/beego"

type TaskController struct {
	beego.Controller
}

// @router /task/add [post]
func (task *TaskController) Add() {

}

// @router /task/:name [delete]
func (task *TaskController) Remove() {

}

// @router /task/:name [put]
func (task *TaskController) Update() {
	taskname := task.Ctx.Input.Param(":name")
	beego.Error("update", taskname)

	task.Data["json"] = "list"
	task.ServeJSON()
}

// @router /task/:name/rename [put]
func (task *TaskController) Rename() {

}

// @router /task/:name [get]
func (task *TaskController) Show() {

}

// @router /task [get]
func (task *TaskController) List() {
	beego.Error("list")

	task.Data["json"] = "list"
	task.ServeJSON()
}
