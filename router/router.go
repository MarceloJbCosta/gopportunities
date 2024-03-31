package router

import (
	"github.com/gin-gonic/gin"
	//o ultimo nome e o que sera refenciado'gin'
)

func Initializer() {
	//inicializa o router utiizando as configuracoes Default do gin
	router := gin.Default()
	//defiinindo um rota usando o verbo GET
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
