package main

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func handle() {
	router := createRouter()
	userRouterGroup(router)
	subjectRouterGroup(router)
	markRouterGroup(router)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	startServer(router, ":8080")
}
