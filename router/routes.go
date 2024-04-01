package router

import (
	"github.com/MarceloJbCosta/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializerRoutes(router *gin.Engine) {
	//initialize Handler

	handler.InitializeHandler()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpeingHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
	}

}
