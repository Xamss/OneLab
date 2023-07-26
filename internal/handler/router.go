package handler

import "github.com/gin-gonic/gin"

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	apiV1 := router.Group("/api/v2/beta")

	user := apiV1.Group("/user")
	user.POST("/register", h.createUser)
	user.POST("/login", h.loginUser)

	post := user.Group("/post")
	post.Use(h.authMiddleware())
	post.GET("/create", h.createArticle)

	return router
}
