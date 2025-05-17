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
	return func(c *gin.Context) {
		var key string
		claims, exists := c.Get(constant.ContextKey.CLAIMS)
		if exists {
			if claims, ok := claims.(utils.Claims); ok {
				key = fmt.Sprintf("rate_limit:%s:%v", c.FullPath(), claims.UserId)
			}
		}
		if key == "" {
			key = fmt.Sprintf("rate_limit:%s:%s", c.FullPath(), c.ClientIP())
		}
		redisClient := config.GetRedisClient()
		count, err := redisClient.Incr(c.Request.Context(), key).Result()
		if err != nil {
			c.Error(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if count == 1 {
			err := redisClient.Expire(c.Request.Context(), key, duration).Err()
			if err != nil {
				c.Error(err)
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
		}
		if count > maxRequest {
			ttl, err := redisClient.TTL(c.Request.Context(), key).Result()
			if err != nil {
				c.Error(err)
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			} else {
				c.Error(errors.New(fmt.Sprintf("Vui lòng thử lại sau %d giây", int(ttl.Seconds()))))
				c.AbortWithStatus(http.StatusTooManyRequests)
				return
			}
		}
		c.Next()
	}
}
