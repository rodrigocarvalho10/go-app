package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigocarvalho10/go-app/schemas"
)

func GetProductionHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	production := schemas.Productions{}
	if err := db.First(&production, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "production not found")
		return
	}

	sendSuccess(ctx, "prodcution found", production)
}
