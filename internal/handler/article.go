package handler

import (
	"api-blog/api"
	"api-blog/internal/entity"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) createArticle(ctx *gin.Context) {

	userID, ok := ctx.MustGet(authUserID).(int64)
	if !ok {
		log.Printf("can't get userID")
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: "can't get user id from auth",
		})
		return
	}

	// logic

	var req entity.Article

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	err = h.srvs.CreateArticle(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.Error{
			Code:    -3,
			Message: err.Error(),
		})
		return
	}

	//logic end

	ctx.JSON(http.StatusCreated, &api.Ok{
		Code:    0,
		Message: "success",
		Data:    userID,
	})
}
