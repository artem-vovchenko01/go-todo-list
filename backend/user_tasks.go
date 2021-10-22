package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	toDoList, err := GetToDoList(c, 422)
	if err != nil {
		SendCustomError(c, 422, "there is no todolist with such id")
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
		SendCustomError(c, 422, "there is no todolist with such id")
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
		SendCustomError(c, 422, "there is no todolist with such id")
		return
	}
	taskId, err := GetTaskId(c)
	if err != nil {
		SendCustomError(c, 401, "taskId parameter invalid")
		return
	}
	_, err = toDoList.Delete(taskId)
	if err != nil {
		SendCustomError(c, http.StatusBadRequest, "there is no task with such id")
		return
	}
	c.Status(204)
}

func UpdateTask(c *gin.Context) {
	toDoList, err := GetToDoList(c, 422)
	if err != nil {
		SendCustomError(c, 422, "there is no todolist with such id")
		return
	}
	task, err := GetTask(c, 422)
	if err != nil {
		SendCustomError(c, 401, "cannot retrieve task")
		return
	}

	newTask := &Task {}
	if err := c.BindJSON(newTask); err != nil {
		SendCustomError(c, 401, "task format invalid")
		return
	}
	newTask.Id = task.Id
	newTask.Status = task.Status
	toDoList.Update(task.Id, newTask)
	c.IndentedJSON(200, newTask)
}
