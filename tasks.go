package main

import (
)

type Task struct {
    Id     int  `json:"id"`
    Name  string  `json:"name"`
    Description string  `json:"description"`
    Status  string `json:"status"`
}

// var Tasks = []Task {
// 	{Id: "10", Name: "Test task", Description: "My cool task", Status: 0},
// 	{Id: "20", Name: "Other task", Description: "My othercool task", Status: 1},
// }


// func GetTasks(c *gin.Context) {
// 	_, claims, err := ParseJWT(c)
// 	if err != nil {
// 		SendError(c, http.StatusBadRequest)
// 	}
// 	userId := UserIds[claims.Email]
// 	todoLists := ToDo[userId]
// 	for _, v := range todoLists {
		
// 	}
// 	c.IndentedJSON(http.StatusOK, Tasks)
// }

// func PostTasks(c *gin.Context) {
// 	var newTask Task

// 	if err := c.BindJSON(&newTask); err != nil {
// 		return
// 	}

// 	Tasks = append(Tasks, newTask)
// 	c.IndentedJSON(http.StatusCreated, newTask)
// }

// func GetTaskById(c *gin.Context) {
// 	id := c.Param("id")
// 	for _, t := range Tasks {
// 		if t.Id == id {
// 			c.IndentedJSON(http.StatusOK, t)
// 			return
// 		}
// 	}

// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "task not found"})
// }
