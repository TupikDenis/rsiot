package main

import (
	"github.com/gin-gonic/gin"
	_ "rsiot/pkg/models"
)

func createRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cORSMiddleware())
	return router
}

func cORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, DELETE, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func userRouterGroup(router *gin.Engine) {
	users := router.Group("/users")
	{
		users.POST("/", addUser)
		users.GET("/", getAllUsers)
		users.GET("/:id", getUsersById)
		users.PATCH("/:id", updateUser)
		users.DELETE("/:id", deleteUser)
	}
}

func subjectRouterGroup(router *gin.Engine) {
	subjects := router.Group("/subjects")
	{
		subjects.POST("/", addSubject)
		subjects.GET("/", getAllSubjects)
		subjects.GET("/:id", getSubjectById)
		subjects.PATCH("/:id", updateSubject)
		subjects.DELETE("/:id", deleteSubject)
	}
}

func markRouterGroup(router *gin.Engine) {
	marks := router.Group("/marks")
	{
		marks.POST("", addMark)
		marks.GET("", getAllMarks)
		marks.GET("/:id", getMarkById)
		marks.PATCH("/:id", updateMark)
		marks.DELETE("/:id", deleteMark)

		marks.GET("/users/:id", getMarkByUserId)
		marks.GET("/subjects/:id", getMarkBySubjectId)
		marks.GET("/users/:id/subjects/:idSub", getMarkByUserAndSubjectId)
	}
}
