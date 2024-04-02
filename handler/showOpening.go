package handler

import (
	"net/http"

	"github.com/MarceloJbCosta/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ShowOpeingHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "Opening not found")
		return
	}

	sendSuccess(ctx, "Show-Openging", opening)

}
