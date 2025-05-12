package middleware

import "github.com/gin-gonic/gin"

func ErrorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()

		if len(context.Errors) > 0 {
			errorMessages := make([]string, len(context.Errors))
			for i, error := range context.Errors {
				errorMessages[i] = error.Error()
			}
			status := context.Writer.Status()
			if len(errorMessages) > 1 {
				context.JSON(status, gin.H{"messages": errorMessages})
			} else {
				context.JSON(status, gin.H{"message": errorMessages[0]})
			}
			context.Abort()
		}
	}
}
