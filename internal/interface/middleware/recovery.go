package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if recovery := recover(); recovery != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("%v", recovery)})
				context.Abort()
			}
		}()
		context.Next()
	}
}
