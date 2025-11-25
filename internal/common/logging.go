package common

import (
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()

		latency := time.Since(t)
		status := c.Writer.Status()

		// Basic access log
		println(
			c.Request.Method,
			c.Request.URL.Path,
			"status=", status,
			"latency=", latency.String(),
		)
	}
}
