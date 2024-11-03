package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	// Initialize router
	router := gin.Default()

	//Initialize router
	InitializeRoutes(router)
	//Run the server port 8080
	router.Run(":8080")
}
