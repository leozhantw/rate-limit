package ratelimit

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/leozhantw/rate-limit/pkg/limiter"
)

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	gin.DefaultWriter = ioutil.Discard

	t.Run("not exceeded", func(t *testing.T) {
		record := limiter.Record{
			Count:     1,
			Remaining: 9,
			ResetAt:   int(time.Now().Unix()) + 1000,
		}
		mock := limiter.NewMockLimiter(ctrl)
		mock.EXPECT().Visit("127.0.0.1").Return(record, nil)
		l = mock

		w := httptest.NewRecorder()
		c, r := gin.CreateTestContext(w)
		r.Use(New())
		r.GET("/foo", func(c *gin.Context) {
			c.String(http.StatusOK, "bar")
		})
		c.Request, _ = http.NewRequest(http.MethodGet, "/foo", nil)
		c.Request.Header.Set("X-Forwarded-For", "127.0.0.1")
		r.ServeHTTP(w, c.Request)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "bar", w.Body.String())
	})

	t.Run("exceeded", func(t *testing.T) {
		record := limiter.Record{
			Count:     10,
			Remaining: 0,
			ResetAt:   int(time.Now().Unix()) + 1000,
		}
		mock := limiter.NewMockLimiter(ctrl)
		mock.EXPECT().Visit("127.0.0.1").Return(record, errors.New("Limit exceeded"))
		l = mock

		w := httptest.NewRecorder()
		c, r := gin.CreateTestContext(w)
		r.Use(New())
		r.GET("/foo", func(c *gin.Context) {
			c.String(http.StatusOK, "bar")
		})
		c.Request, _ = http.NewRequest(http.MethodGet, "/foo", nil)
		c.Request.Header.Set("X-Forwarded-For", "127.0.0.1")
		r.ServeHTTP(w, c.Request)

		assert.Equal(t, http.StatusTooManyRequests, w.Code)
		assert.Equal(t, "Error", w.Body.String())
	})
}
