package routes

import (
	"task-api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incoming *gin.Engine) {
	incoming.GET("/users/:id", controllers.GetUser())
	incoming.GET("/users", controllers.GetUsers())
}
