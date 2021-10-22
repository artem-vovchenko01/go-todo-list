package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetCurrentUserToDo(c *gin.Context) (*UserToDo, error) {
	_, claims, err := ParseJWT(c)
	if err != nil {
		SendError(c, http.StatusBadRequest)
	}
	return Storage.GetUserToDoByEmail(claims.Email)
}

func GetToDoList(c *gin.Context, customErrorCode int) (*ToDoList, error) {
	userToDo, err := GetCurrentUserToDo(c)
	if err != nil {
		SendCustomError(c, customErrorCode, "there is no todo list for this user")
	}

	listIdStr := c.Param("listId")
	listId, err := strconv.Atoi(listIdStr)
	if err != nil {
		SendCustomError(c, customErrorCode, "listId parameter invalid")
		return nil, err
	}
	return userToDo.Get(listId)
}

func GetTaskId(c *gin.Context) (int, error) {
	taskIdStr := c.Param("taskId")
	return strconv.Atoi(taskIdStr)
}

func GetTask(c *gin.Context, customErrorCode int) (*Task, error) {
	toDoList, err := GetToDoList(c, customErrorCode)
	if err != nil {
		SendCustomError(c, http.StatusBadRequest, "error while retrieving todo list")
		return nil, err
	}
	taskId, err := GetTaskId(c)
	if err != nil {
		SendCustomError(c, customErrorCode, "error while retrieving taskId")
	}
	return toDoList.Get(taskId)
}
