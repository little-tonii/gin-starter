package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if recovery := recover(); recovery != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("%v", recovery)})
				c.Abort()
			}
		}()
		c.Next()
	}
}
