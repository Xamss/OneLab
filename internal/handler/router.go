package handler

import (
	_ "api-blog/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiV1 := router.Group("/api/v1")

	user := apiV1.Group("/user")
	user.POST("/register", h.createUser)
	user.POST("/login", h.loginUser)

	post := user.Group("/post")
	post.Use(h.authMiddleware())
	post.GET("/create", h.createArticle)

	return router
}
