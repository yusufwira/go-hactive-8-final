package handler

import (
	"final-project/service/mygram/dto"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h MyGramHandler) Register(ctx *gin.Context) {
	req := dto.RequestReqister{}
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Errors",
			"data":    err.Error(),
		})
		return
	}
	isvalid := req.ValidateRegister()
	if isvalid != nil {
		errRes := strings.Split(isvalid.Error(), ";")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid Request",
			"data":    errRes,
		})
		return
	}

	res, err := h.MyGramUsecase.RegisterUser(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Error",
			"data":    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
		"data":    res,
	})
}

func (h MyGramHandler) Login(ctx *gin.Context) {
	req := dto.RequestLogin{}
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Errors",
			"data":    err.Error(),
		})
		return
	}
	isvalid := req.Validate()
	if isvalid != nil {
		errRes := strings.Split(isvalid.Error(), ";")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid Request",
			"data":    errRes,
		})
		return
	}

	res, err := h.MyGramUsecase.LoginUser(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Error",
			"data":    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
		"data":    res,
	})
}
