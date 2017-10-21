package controllers

import (
	"encoding/json"
	"sir/lib/config"
	"sir/models"

	"github.com/astaxie/beego"
	"sir/task"
)

var TaskManager *task.TaskManager

func init() {
	TaskManager = task.NewTaskManager(config.AppHome + beego.AppConfig.String("workspace"))
}

type TaskController struct {
	beego.Controller
}

// @router /task/add [post]
func (task *TaskController) Add() {
	var ob models.TaskConfig
	json.Unmarshal(task.Ctx.Input.RequestBody, &ob)

	config.CreateTaskConfig(&ob)
	// TODO handle error

	task.Data["json"] = map[string]interface{}{"data": ob}
	task.ServeJSON()
}

// @router /task/:name [delete]
func (task *TaskController) Remove() {
	taskname := task.Ctx.Input.Param(":name")
	config.DeleteTaskConfig(taskname)
	// TODO handle error

	// TODO update mem

	task.ServeJSON()
}

// @router /task/:name [put]
func (task *TaskController) Update() {
	taskname := task.Ctx.Input.Param(":name")
	// TODO just open editor

	task.Data["json"] = map[string]interface{}{"data": config.GetTaskConfigFilePath(taskname)}

	task.ServeJSON()
}

// @router /task/:name/rename [put]
func (task *TaskController) Rename() {
}

// @router /task/:name [get]
func (task *TaskController) Show() {
	taskname := task.Ctx.Input.Param(":name")

	taskConfig, _ := config.GetTaskConfig(taskname)
	// TODO handle error

	// TODO get task state info

	task.Data["json"] = map[string]interface{}{"data": *taskConfig}
	task.ServeJSON()
}

// @router /task [get]
func (task *TaskController) List() {
	// beego.Error("list")

	taskConfigs := config.ListAllTaskConfigs()

	tasks := make([]models.Task, 0, 10)
	for _, c := range taskConfigs {
		tasks = append(tasks, models.Task{
			TaskConfig: &c,
		})
	}

	// TODO add task state

	task.Data["json"] = map[string][]models.Task{
		"data": tasks,
	}
	task.ServeJSON()
}
