package handler

import (
	"api-blog/internal/entity"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) createArticle(ctx *gin.Context) {
	var req entity.Article

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &Error{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}

	err = h.srvs.CreateArticle(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	ctx.Status(http.StatusCreated)
}
