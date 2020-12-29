package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(app *gin.Engine) {
	app.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "")
	})
}
