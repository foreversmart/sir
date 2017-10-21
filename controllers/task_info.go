package controllers

// @router /task/:name/statistics [get]
func (task *TaskController) Statistics() {

}

// @router /task/:name/log [get]
func (task *TaskController) Log() {
	taskname := task.Ctx.Input.Param(":name")
	// TODO just open editor

	println(taskname)
	task.Data["json"] = map[string]interface{}{"data": "/Users/alex/Desktop/log.log"}

	task.ServeJSON()
}

// @router /task/log [get]
func (task *TaskController) AllLog() {
	task.Data["json"] = map[string]interface{}{"data": "/Users/alex/Desktop/log.log"}

	task.ServeJSON()
}
