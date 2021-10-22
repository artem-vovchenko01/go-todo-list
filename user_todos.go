package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateToDoList(c *gin.Context) {
	_, claims, err := ParseJWT(c)
	if err != nil {
		SendError(c, http.StatusBadRequest)
	}
	userToDo, err := Storage.GetOrCreateUserToDoByEmail(claims.Email)
	if err != nil {
		SendCustomError(c, http.StatusBadRequest, "unexpected server error")
	}

	toDoList := NewToDoListStorage()
	c.BindJSON(toDoList)

	err = userToDo.Add(toDoList)
	c.IndentedJSON(201, toDoList)
}

func UpdateToDoList(c *gin.Context) {
	_, claims, err := ParseJWT(c)
	if err != nil {
		SendError(c, http.StatusBadRequest)
	}
	userToDo, err := Storage.GetUserToDoByEmail(claims.Email)
	if err != nil {
		SendCustomError(c, http.StatusBadRequest, "there is no todo list for this user")
	}
	listIdStr := c.Param("listId")
	var listId int
	if listId, err = strconv.Atoi(listIdStr); err != nil {
		SendCustomError(c, http.StatusBadRequest, "listId parameter invalid")
		return
	}
	toDoList, err := userToDo.Get(listId)
	if err != nil {
		SendCustomError(c, http.StatusBadRequest, "there is no todolist with such id")
		return
	}

	var newToDoList *ToDoList = &ToDoList{}
	c.BindJSON(newToDoList)
	newToDoList.tasks = toDoList.tasks
	newToDoList.currentTaskId = toDoList.currentTaskId
	newToDoList.Id = toDoList.Id
	newToDoList.lock = toDoList.lock
	userToDo.Update(listId, newToDoList)
	c.IndentedJSON(204, newToDoList)
}

func GetToDoLists(c *gin.Context) {
	_, claims, err := ParseJWT(c)
	if err != nil {
		SendError(c, http.StatusBadRequest)
	}

	userToDo, err := Storage.GetUserToDoByEmail(claims.Email)
	if err != nil {
		c.IndentedJSON(200, "{}")
		return
	}
	listsSlice := make([]*ToDoList, 0, len(userToDo.lists))
	for _, v := range userToDo.lists {
		listsSlice = append(listsSlice, v)
	}
	c.IndentedJSON(200, listsSlice)
}
