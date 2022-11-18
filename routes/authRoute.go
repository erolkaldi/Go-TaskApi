package routes

import (
	"task-api/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incoming *gin.Engine) {
	incoming.POST("auth/register", controllers.RegisterUser())
	incoming.POST("auth/login", controllers.GetToken())

}
