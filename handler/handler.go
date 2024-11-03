package handler

import (
	"github.com/rodrigocarvalho10/go-app/configuration"
	"gorm.io/gorm"
)

var (
	logger *configuration.Logger
	db     *gorm.DB
)

func InitializerHandler() {
	logger = configuration.GetLogger("handler")
	db = configuration.GetSQLite()
}
