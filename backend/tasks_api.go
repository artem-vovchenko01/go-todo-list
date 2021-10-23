package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	toDoList, err := GetToDoList(c, 422)
	if err != nil {
		SendCustomError(c, 422, errNoToDoListWithSuchId)
		return
	}

	var task *Task = &Task{}
	c.BindJSON(task)
	toDoList.Add(task)
	c.IndentedJSON(201, task)
}

func GetTasks(c *gin.Context) {
	toDoList, err := GetToDoList(c, 422)
	if err != nil {
		SendCustomError(c, 422, errNoToDoListWithSuchId)
		return
	}

	taskSlice := make([]*Task, 0, len(toDoList.tasks))
	for _, v := range toDoList.tasks {
		taskSlice = append(taskSlice, v)
	}
	c.IndentedJSON(200, taskSlice)
}

func DeleteTask(c *gin.Context) {
	toDoList, err := GetToDoList(c, 422)
	if err != nil {
		SendCustomError(c, 422, errNoToDoListWithSuchId)
		return
	}
	taskId, err := GetTaskId(c)
	if err != nil {
		SendCustomError(c, 401, errTaskIdParamInvalid)
		return
	}
	_, err = toDoList.Delete(taskId)
	if err != nil {
		SendCustomError(c, http.StatusBadRequest, errNoTaskWithSuchId)
		return
	}
	c.Status(204)
}

func UpdateTask(c *gin.Context) {
	toDoList, err := GetToDoList(c, 422)
	if err != nil {
		SendCustomError(c, 422, errNoToDoListWithSuchId)
		return
	}
	task, err := GetTask(c, 422)
	if err != nil {
		SendCustomError(c, 401, errCannotRetrieveTask)
		return
	}

	newTask := &Task{}
	newTask.Status = task.Status
	newTask.Description = task.Description
	newTask.Name = task.Name
	if err := c.BindJSON(newTask); err != nil {
		SendCustomError(c, 401, errTaskFormatInvalid)
		return
	}
	newTask.Id = task.Id
	toDoList.Update(task.Id, newTask)
	c.IndentedJSON(200, newTask)
}
