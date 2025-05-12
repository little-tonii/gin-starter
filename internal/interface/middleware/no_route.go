package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NoRoute() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{"message": "Điểm truy cập không tồn tại"})
	}
}
