package handler

import (
	"final-project/service/mygram/dto"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (h MyGramHandler) PostComment(ctx *gin.Context) {
	req := dto.RequestPostComment{}
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	userStr := strconv.FormatUint(uint64(userId), 10)

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

	res, err := h.MyGramUsecase.PostComment(ctx, req, userStr)
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

func (h MyGramHandler) UpdateComment(ctx *gin.Context) {
	req := dto.RequestPostComment{}
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	userStr := strconv.FormatUint(uint64(userId), 10)
	ids := ctx.Param("id")
	if ids == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Errors",
		})
		return
	}

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

	_, err := h.MyGramUsecase.UpdateComment(ctx, ids, userStr, req)
	if err != nil {
		ctx.JSON(http.StatusNonAuthoritativeInfo, gin.H{
			"status":  http.StatusNonAuthoritativeInfo,
			"message": "Not Authorize",
			"data":    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
	})
}

func (h MyGramHandler) DeleteComment(ctx *gin.Context) {
	ids := ctx.Param("id")
	if ids == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Errors",
		})
		return
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	userStr := strconv.FormatUint(uint64(userId), 10)

	_, err := h.MyGramUsecase.DeleteComment(ctx, ids, userStr)
	if err != nil {
		ctx.JSON(http.StatusNonAuthoritativeInfo, gin.H{
			"status":  http.StatusNonAuthoritativeInfo,
			"message": "Not Authorize",
			"data":    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
	})
}

func (h MyGramHandler) GetAllComment(ctx *gin.Context) {
	res, err := h.MyGramUsecase.GetAllComment(ctx)
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

func (h MyGramHandler) GetOneComment(ctx *gin.Context) {
	ids := ctx.Param("id")
	if ids == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Errors",
		})
		return
	}

	res, err := h.MyGramUsecase.GetOneComment(ctx, ids)
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
