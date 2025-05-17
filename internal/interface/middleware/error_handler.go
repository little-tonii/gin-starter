package middleware

import "github.com/gin-gonic/gin"

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			errorMessages := make([]string, len(c.Errors))
			for i, error := range c.Errors {
				errorMessages[i] = error.Error()
			}
			status := c.Writer.Status()
			if len(errorMessages) > 1 {
				c.JSON(status, gin.H{"messages": errorMessages})
			} else {
				c.JSON(status, gin.H{"message": errorMessages[0]})
			}
			c.Abort()
		}
	}
}
