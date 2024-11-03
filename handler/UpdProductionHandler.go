package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigocarvalho10/go-app/schemas"
)

//	@BasePath	/api/v1

//	@Summary		Update production
//	@Description	Update a movie or serie
//	@Tags			productions
//	@Accept			json
//	@Produce		json
//	@Param			id		query		string					true	"production Identification"
//	@Param			production	body		UpdateproductionRequest	true	"production data to Update"
//	@Success		200		{object}	UpdateproductionResponse
//	@Failure		400		{object}	ErrorResponse
//	@Failure		404		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/production [put]

func UpdProductionHandler(ctx *gin.Context) {
	request := UpdateProductionRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

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
	// Update production
	if request.Name != "" {
		production.Name = request.Name
	}

	if request.Producer != "" {
		production.Producer = request.Producer
	}

	if request.Movie != nil {
		production.Movie = *request.Movie
	}

	if request.Series != nil {
		production.Series = *request.Series
	}

	if request.Protagonist != "" {
		production.Protagonist = request.Protagonist
	}

	if request.Notice > 0 {
		production.Notice = request.Notice
	}

	if request.Assessment != "" {
		production.Assessment = request.Assessment
	}

	// Save production
	if err := db.Save(&production).Error; err != nil {
		logger.Errorf("error updating production: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating production")
		return
	}
	sendSuccess(ctx, "update-production", production)
}
