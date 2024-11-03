package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigocarvalho10/go-app/schemas"
)

//	@BasePath	/api/v1

//	@Summary		Delete production
//	@Description	Delete a new job production
//	@Tags			productions
//	@Accept			json
//	@Produce		json
//	@Param			id	query		string	true	"production identification"
//	@Success		200	{object}	DeleteproductionResponse
//	@Failure		400	{object}	ErrorResponse
//	@Failure		404	{object}	ErrorResponse
//	@Router			/handler [delete]

func DeleteProductionHandler(ctx *gin.Context) {

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	production := schemas.Productions{}

	if err := db.First(&production, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("Productions with id:  %s not found", id))
		return
	}
	if err := db.Delete(&production, id).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting production with id: %s", id))
		return
	}

	sendSuccess(ctx, "delete production", production)
}
