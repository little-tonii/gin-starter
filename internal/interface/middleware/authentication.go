package middleware

import (
	"errors"

	"health-care-system/internal/infrastructure/utils"
	"health-care-system/internal/shared/constant"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.Request.Header.Get("Authorization")
		if authHeader == "" {
			context.Error(errors.New("Người dùng chưa đăng nhập"))
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		const prefix = "Bearer "
		if len(authHeader) <= len(prefix) || authHeader[:len(prefix)] != prefix {
			context.Error(errors.New("Authorization header không hợp lệ"))
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := authHeader[len(prefix):]

		claims, err := utils.VerifyToken(constant.Environment.JWT_SECRET_KEY, token)

		if err != nil {
			context.Error(errors.New(err.Error()))
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		context.Set("claims", claims)
		context.Next()
	}
}
