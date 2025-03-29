package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func makeHelloWorld(opt string) string {
	return fmt.Sprintf("Hello World! %s", opt)
}

func main() {
	router := gin.Default()
	router.GET("/:opt", func(ctx *gin.Context) {
		opt := ctx.Param("opt")
		ctx.String(http.StatusOK, makeHelloWorld(opt))
	})
	router.Run(":8080")
}
