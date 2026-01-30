package router

import (
	"github.com/gin-gonic/gin"
	docs "github.com/rodrigocarvalho10/go-app/docs"
	"github.com/rodrigocarvalho10/go-app/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(router *gin.Engine) {
	//Initializer handler
	handler.InitializerHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET("/production", handler.GetProductionHandler)
		v1.POST("/production", handler.PostProductionHandler)
		v1.DELETE("/rmproduction", handler.DeleteProductionHandler)
		v1.PUT("/updproduction", handler.UpdProductionHandler)
		v1.GET("/productions", handler.LisProductionsHandler)
	}
	//Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
