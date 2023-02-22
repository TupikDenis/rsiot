package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func startServer(router *gin.Engine, port string) {
	router.Use(cors.Default())

	err := router.Run(port)

	if err != nil {
		panic(err)
	}
}
