package ratelimit

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/leozhantw/rate-limit/pkg/limiter"
)

const (
	max = 10
	ttl = 60
)

var l limiter.Limiter

func New() gin.HandlerFunc {
	if l == nil {
		l = limiter.New(max, ttl)
	}

	return func(c *gin.Context) {
		record, err := l.Visit(c.ClientIP())
		if err != nil {
			c.String(http.StatusTooManyRequests, "Error")
			c.Abort()
			return
		}

		c.Writer.Header().Set("X-RateLimit-Limit", strconv.Itoa(max))
		c.Writer.Header().Set("X-RateLimit-Remaining", strconv.Itoa(record.Remaining))
		c.Writer.Header().Set("X-RateLimit-Reset", strconv.Itoa(record.ResetAt))
	}
}
