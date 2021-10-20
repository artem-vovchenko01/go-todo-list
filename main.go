package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/tasks", getTasks)

	router.Run("localhost:8080")
}

// album represents data about a record album.
type Task struct {
    Id     string  `json:"id"`
    Name  string  `json:"name"`
    Description string  `json:"description"`
    Status  float64 `json:"status"`
}

var tasks = []Task {
	{Id: "10", Name: "Test task", Description: "My cool task", Status: 0},
	{Id: "20", Name: "Other task", Description: "My othercool task", Status: 1},
}

func getTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}
