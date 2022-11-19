package routes

import (
	"task-api/controllers"
	"task-api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incoming *gin.Engine) {
	incoming.Use(middleware.Authenticate())
	incoming.GET("/users/:id", controllers.GetUser())
	incoming.GET("/users", controllers.GetUsers())
	incoming.POST("/users/changepassword", controllers.ChangePassword())
}
