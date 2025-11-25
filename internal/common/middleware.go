package common

import "github.com/gin-gonic/gin"

func GetRestaurantID(c *gin.Context) int {
	id := c.GetInt("restaurant_id")
	return id
}

func GetUserEmail(c *gin.Context) string {
	return c.GetString("email")
}

func GetUserID(c *gin.Context) uint {
	val, ok := c.Get("user_id")
	if !ok {
		return 0
	}
	return val.(uint)
}
