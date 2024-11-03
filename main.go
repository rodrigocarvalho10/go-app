package main

import (
	"github.com/rodrigocarvalho10/go-app/configuration"
	"github.com/rodrigocarvalho10/go-app/router"
)

var (
	logger configuration.Logger
)

//	@title			Movies and Series - Productions
//	@version		0.1.2
//  @BasePath		/api/v1
//  @Host			github.com/rodrigocarvalho10/go-app
//	@description	This is an example of an API for registering film productions
//	@termsOfService	http://swagger.io/terms/

func main() {
	logger = *configuration.GetLogger("main")
	//Initialize Configs
	err := configuration.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}
	// Initialize router
	router.Initialize()
}
