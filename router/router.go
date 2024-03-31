package router

import (
	"github.com/gin-gonic/gin"
	//o ultimo nome e o que sera refenciado'gin'
)

func Initializer() {

	//initialize Router

	router := gin.Default()

	//initializer routes

	initializerRoutes(router)

	//run the server
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
