package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

var errLimitExceed = errors.New("Rate limit exceed!")

func Limit(bkt *ratelimit.Bucket) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if bkt.TakeAvailable(1) == 0 {
			ctx.JSON(http.StatusServiceUnavailable, gin.H{
				"code":   503,
				"msg":    "失败",
				"errors": []string{errLimitExceed.Error()},
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
