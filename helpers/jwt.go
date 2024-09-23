package helpers

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func VerifyToken(ctx *gin.Context) (interface{}, error) {
	errResponse := errors.New("sign in process failed")
	headerToken := ctx.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer ")

	if !bearer {
		return nil, errResponse
	}

	stringToken := strings.TrimPrefix(headerToken, "Bearer ")
	token, err := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, errResponse
		}
		return &privateKey.PublicKey, nil
	})

	if err != nil {
		return nil, errResponse
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errResponse
}
