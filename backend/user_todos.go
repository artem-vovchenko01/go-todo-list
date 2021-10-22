package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	userToDo, err := GetCurrentUserToDo(c)
	if err != nil {
		SendCustomError(c, http.StatusBadRequest, "error while retrieving usert todo")
		return
	}

	toDoList, err := GetToDoList(c, http.StatusBadRequest)
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
	userToDo.Update(toDoList.Id, newToDoList)
	c.IndentedJSON(204, newToDoList)
}

func GetToDoLists(c *gin.Context) {
	userToDo, err := GetCurrentUserToDo(c)
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

func DeleteToDoListHelper(c *gin.Context, listId int) error {
	_, claims, err := ParseJWT(c)
	if err != nil {
		SendError(c, http.StatusBadRequest)
		return err
	}

	userToDo, err := Storage.GetUserToDoByEmail(claims.Email)
	if err != nil {
		SendError(c, http.StatusBadRequest)
		return err
	}
	userToDo.Delete(listId)
	return nil
}

func DeleteToDoList(c *gin.Context) {
	toDoList, err := GetToDoList(c, 422)
	if err != nil {
		SendCustomError(c, 422, "there is no todolist with such id")
		return
	}

	for k := range toDoList.tasks {
		toDoList.Delete(k)
	}

	DeleteToDoListHelper(c, toDoList.Id)
	c.Status(204)
}
