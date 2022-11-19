package main

import (
	"os"
	routes "task-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	router := gin.Default()
	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	routes.DefinitionRoutes(router)
	router.Run(port)

}
