package middleware

import (
	"errors"

	"gin-starter/internal/shared/constant"
	"gin-starter/internal/shared/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.Error(errors.New("Người dùng chưa đăng nhập"))
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		const prefix = "Bearer "
		if len(authHeader) <= len(prefix) || authHeader[:len(prefix)] != prefix {
			c.Error(errors.New("Authorization header không hợp lệ"))
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := authHeader[len(prefix):]

		claims, err := utils.VerifyToken(constant.Environment.JWT_SECRET_KEY, token)

		if err != nil {
			c.Error(errors.New(err.Error()))
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set(constant.ContextKey.CLAIMS, claims)
		c.Next()
	}
}
