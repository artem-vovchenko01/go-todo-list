package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var Storage StorageService = *NewStorageService()

func main() {
	router := gin.Default()

	router.POST("/user/signin", Signin)
	router.POST("/user/signup", Signup)

	userRoutes := router.Group("/user", AuthorizeJWT()) 
	{
		// userRoutes.GET("/tasks", GetTasks)
		// userRoutes.GET("/tasks/:id", GetTaskById)
		userRoutes.GET("/welcome", Welcome)
		userRoutes.GET("/jwt/refresh", Refresh)

		// userRoutes.POST("/tasks", PostTasks)
	}

	todoRoutes := router.Group("/todo", AuthorizeJWT()) 
	{
		todoRoutes.POST("/lists", CreateToDoList)
		todoRoutes.PUT("/lists/:listId", UpdateToDoList)
		todoRoutes.GET("/lists", GetToDoLists)

		todoRoutes.POST("/lists/:listId/tasks", CreateTask)
		todoRoutes.GET("/lists/:listId/tasks", GetTasks)
	}

	router.Run("localhost:8080")
}

func Welcome(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H {"message" : "welcome"})
}

func SendResponse(c *gin.Context, status int, resp string) {
	// c.IndentedJSON(status, gin.H { "message" : httpStatusMessages[status] })
	c.Writer.Write([]byte(resp))
	c.Status(status)
}

// func SendResponseJSON(c *gin.Context, status int, gin.H)

func SendError(c *gin.Context, status int) {
	c.IndentedJSON(status, gin.H { "message" : httpStatusMessages[status] })
}

func SendCustomError(c *gin.Context, status int, message string) {
	c.IndentedJSON(status, gin.H { "message" : message })
}
