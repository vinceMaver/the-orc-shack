package common

import (
	"net/http"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gin-gonic/gin"
)

func RateLimitMiddleware(limit int) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.GetString("email")
		if email == "" {
			c.Next()
			return
		}

		key := "rate:" + email

		// Try incrementing
		newCount, err := Cache.Increment(key, 1)

		if err == memcache.ErrCacheMiss {
			// Key doesn't exist → create with value 1 and TTL
			err = Cache.Set(&memcache.Item{
				Key:        key,
				Value:      []byte("1"),
				Expiration: int32(time.Minute.Seconds()),
			})
			if err != nil {
				// allow request on cache failure
				c.Next()
				return
			}
			newCount = 1
		}

		if err != nil {
			// Allow request if Memcached fails
			c.Next()
			return
		}

		if int(newCount) > limit {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
