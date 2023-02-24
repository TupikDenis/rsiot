package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func startServer(router *gin.Engine, port string) {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	err := router.Run(port)

	if err != nil {
		panic(err)
	}
}
