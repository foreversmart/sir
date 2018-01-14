package controllers

import (
	"sir/lib/config"
	"sir/models"

	"github.com/astaxie/beego"
)

// @router /task/:name/start [post]
func (task *TaskController) Start() {
	taskname := task.Ctx.Input.Param(":name")
	if taskname == "" {
		beego.Error("should have task name ")
		return
	}

	conf, err := config.GetTaskConfig(taskname)
	if err != nil {
		beego.Error("config.GetTaskConfig, ", err)
		return
	}

	err = TaskManager.StartTask(&models.Task{
		TaskConfig: conf,
	})
	if err != nil {
		beego.Error("TaskManager.StartTask", err)
		return
	}

	task.ServeJSON()
}

// @router /task/:name/restart [post]
func (task *TaskController) Restart() {
	taskname := task.Ctx.Input.Param(":name")
	if taskname == "" {
		beego.Error("should have task name ")
		return
	}

	err := TaskManager.RemoveTask(taskname)
	if err != nil {
		beego.Error("TaskManager.RemoveTask", err)
		return
	}

	conf, err := config.GetTaskConfig(taskname)
	if err != nil {
		beego.Error("config.GetTaskConfig, ", err)
		return
	}

	err = TaskManager.StartTask(&models.Task{
		TaskConfig: conf,
	})
	if err != nil {
		beego.Error("TaskManager.StartTask", err)
		return
	}

	task.ServeJSON()
}

// @router /task/:name/stop [post]
func (task *TaskController) Stop() {
	taskname := task.Ctx.Input.Param(":name")
	if taskname == "" {
		beego.Error("should have task name ")
		return
	}

	err := TaskManager.RemoveTask(taskname)
	if err != nil {
		beego.Error("TaskManager.RemoveTask", err)
		return
	}

	task.ServeJSON()
}

// @router /task/:name/send [post]
func (task *TaskController) Send() {
	// TODO
}
