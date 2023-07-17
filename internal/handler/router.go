package handler

import "github.com/gin-gonic/gin"

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	apiV1 := router.Group("/api/v2/beta")

	apiV1.POST("/user-register", h.createUser)
	apiV1.POST("/article-create", h.createArticle)

	return router
}
