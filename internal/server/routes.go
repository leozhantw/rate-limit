package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/leozhantw/rate-limit/internal/middleware/ratelimit"
)

func Routes(app *gin.Engine) {
	app.Use(ratelimit.New())

	app.GET("/", func(c *gin.Context) {
		max, _ := strconv.Atoi(c.Writer.Header().Get("X-RateLimit-Limit"))
		remaining, _ := strconv.Atoi(c.Writer.Header().Get("X-RateLimit-Remaining"))
		hit := max - remaining

		c.String(http.StatusOK, strconv.Itoa(hit))
	})
}
