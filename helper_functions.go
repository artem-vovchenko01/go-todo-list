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

func GetToDoList(c *gin.Context, userToDo *UserToDo, customErrorCode int) (*ToDoList, error) {
	listIdStr := c.Param("listId")
	listId, err := strconv.Atoi(listIdStr)
	if err != nil {
		SendCustomError(c, customErrorCode, "listId parameter invalid")
		return nil, err
	}
	return userToDo.Get(listId)
}
