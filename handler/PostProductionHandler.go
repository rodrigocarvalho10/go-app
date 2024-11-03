package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigocarvalho10/go-app/schemas"
)

func PostProductionHandler(ctx *gin.Context) {

	request := CreateProductionRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	production := schemas.Productions{
		Name:        request.Name,
		Producer:    request.Producer,
		Movie:       *request.Movie,
		Series:      *request.Series,
		Protagonist: request.Protagonist,
		Notice:      request.Notice,
		Assessment:  request.Assessment,
	}

	if err := db.Create(&production).Error; err != nil {
		logger.Errorf("error creating production: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating production on database")
		return
	}

	sendSuccess(ctx, "create-production", production)
}
