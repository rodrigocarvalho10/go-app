package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigocarvalho10/go-app/schemas"
)

func LisProductionsHandler(ctx *gin.Context) {
	productions := []schemas.Productions{}

	if err := db.Find(&productions).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing productions")
		return
	}

	sendSuccess(ctx, "list-productions", productions)
}
