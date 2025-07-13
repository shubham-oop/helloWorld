package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shubham-oop/helloWorld/internal/handler"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/students", handler.GetStudents)
		api.GET("/students/:roll_number", handler.GetStudentByRollNumber)
		api.POST("/students", handler.AddStudent)

	}

	return r
}
