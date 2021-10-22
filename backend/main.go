package main

import (
	"github.com/gin-gonic/gin"
)

var Storage StorageService = *NewStorageService()

func main() {
	router := gin.Default()

	router.POST("/user/signin", Signin)
	router.POST("/user/signup", Signup)

	userRoutes := router.Group("/user", AuthorizeJWT()) 
	{
		userRoutes.GET("/jwt/refresh", Refresh)
	}

	todoRoutes := router.Group("/todo", AuthorizeJWT()) 
	{
		todoRoutes.GET("/lists", GetToDoLists)
		todoRoutes.POST("/lists", CreateToDoList)
		todoRoutes.PUT("/lists/:listId", UpdateToDoList)
		todoRoutes.DELETE("/lists/:listId", DeleteToDoList)

		todoRoutes.GET("/lists/:listId/tasks", GetTasks)
		todoRoutes.POST("/lists/:listId/tasks", CreateTask)
		todoRoutes.PUT("/lists/:listId/tasks/:taskId", UpdateTask)
		todoRoutes.DELETE("/lists/:listId/tasks/:taskId", DeleteTask)
	}

	router.Run("localhost:8080")
}

func SendResponse(c *gin.Context, status int, resp string) {
	c.Writer.Write([]byte(resp))
	c.Status(status)
}

func SendError(c *gin.Context, status int) {
	c.IndentedJSON(status, gin.H { "message" : httpStatusMessages[status] })
}

func SendCustomError(c *gin.Context, status int, message string) {
	c.IndentedJSON(status, gin.H { "message" : message })
}
