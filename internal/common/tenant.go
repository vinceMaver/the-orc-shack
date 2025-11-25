package common

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TenantMiddleware ensures each request belongs to a restaurant (SaaS tenant)
func TenantMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		raw := c.GetHeader("X-Restaurant-ID")
		if raw == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Missing X-Restaurant-ID header",
			})
			c.Abort()
			return
		}

		id, err := strconv.Atoi(raw)
		if err != nil || id <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid restaurant ID",
			})
			c.Abort()
			return
		}

		c.Set("restaurant_id", id)
		c.Next()
	}
}
