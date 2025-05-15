package middleware

import (
	"errors"
	"fmt"
	"gin-starter/internal/infrastructure/config"
	"gin-starter/internal/shared/constant"
	"gin-starter/internal/shared/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RateLimit(maxRequest int64, duration time.Duration) gin.HandlerFunc {
	return func(context *gin.Context) {
		var key string
		claims, exists := context.Get(constant.ContextKey.CLAIMS)
		if exists {
			if claims, ok := claims.(utils.Claims); ok {
				key = fmt.Sprintf("rate_limit:%s:%v", context.FullPath(), claims.UserId)
			}
		}
		if key == "" {
			key = fmt.Sprintf("rate_limit:%s:%s", context.FullPath(), context.ClientIP())
		}
		redisClient := config.GetRedisClient()
		count, err := redisClient.Incr(context.Request.Context(), key).Result()
		if err != nil {
			context.Error(err)
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if count == 1 {
			err := redisClient.Expire(context.Request.Context(), key, duration).Err()
			if err != nil {
				context.Error(err)
				context.AbortWithStatus(http.StatusInternalServerError)
				return
			}
		}
		if count > maxRequest {
			ttl, err := redisClient.TTL(context.Request.Context(), key).Result()
			if err != nil {
				context.Error(err)
				context.AbortWithStatus(http.StatusInternalServerError)
				return
			} else {
				context.Error(errors.New(fmt.Sprintf("Vui lòng thử lại sau %d giây", int(ttl.Seconds()))))
				context.AbortWithStatus(http.StatusTooManyRequests)
				return
			}
		}
		context.Next()
	}
}
