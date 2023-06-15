// ./main.go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main()  {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) { // route home url
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello Word!",
		})
	})

	r.Run(":5050")
}