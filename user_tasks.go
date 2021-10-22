package main

import (
	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	userToDo, err := GetCurrentUserToDo(c)
	if err != nil {
		SendCustomError(c, 422, "there is no todo list for this user")
		return
	}
	toDoList, err := GetToDoList(c, userToDo, 422)
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
	userToDo, err := GetCurrentUserToDo(c)
	if err != nil {
		SendCustomError(c, 422, "there is no todo list for this user")
		return
	}
	toDoList, err := GetToDoList(c, userToDo, 422)
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
