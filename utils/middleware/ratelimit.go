package middleware

import (
	"CSAwork/global"
	"CSAwork/utils"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

func RateLimitMiddleware(fillInterval time.Duration, cap int64) func(c *gin.Context) {
	bucket := ratelimit.NewBucket(fillInterval, cap)
	global.Bucket = bucket
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			utils.RespFail(c, "rate limit")
			c.Abort()
			return
		}
		c.Next()
	}
}
