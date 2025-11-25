package restaurant

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateRestaurantBody struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

func CreateRestaurantHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body CreateRestaurantBody
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		r := Restaurant{
			Name:        body.Name,
			Description: body.Description,
		}

		db.Create(&r)
		c.JSON(201, r)
	}
}

func ListRestaurantsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var list []Restaurant
		db.Find(&list)
		c.JSON(200, list)
	}
}
