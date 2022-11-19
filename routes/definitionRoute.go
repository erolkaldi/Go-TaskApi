package routes

import (
	"task-api/controllers"
	"task-api/middleware"

	"github.com/gin-gonic/gin"
)

func DefinitionRoutes(incoming *gin.Engine) {
	incoming.Use(middleware.Authenticate())
	incoming.POST("definition/addtasktype", controllers.AddTaskType())

}
