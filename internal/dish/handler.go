package dish

import (
	"orc-shack/internal/common"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DishBody struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"image_url"`
}

func CreateDishHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body DishBody
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		rid := common.GetRestaurantID(c)

		d := Dish{
			Name:         body.Name,
			Description:  body.Description,
			Price:        body.Price,
			ImageURL:     body.ImageURL,
			RestaurantID: uint(rid),
		}

		db.Create(&d)
		c.JSON(201, d)
	}
}

func ListDishesHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rid := common.GetRestaurantID(c)

		var list []Dish
		db.Where("restaurant_id = ?", rid).Find(&list)

		c.JSON(200, list)
	}
}

func GetDishHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rid := common.GetRestaurantID(c)
		id := c.Param("id")

		var d Dish
		if db.Where("id = ? AND restaurant_id = ?", id, rid).First(&d).Error != nil {
			c.JSON(404, gin.H{"error": "not found"})
			return
		}

		c.JSON(200, d)
	}
}

func UpdateDishHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rid := common.GetRestaurantID(c)
		id := c.Param("id")

		var existing Dish
		if db.Where("id = ? AND restaurant_id = ?", id, rid).First(&existing).Error != nil {
			c.JSON(404, gin.H{"error": "not found"})
			return
		}

		var body DishBody
		c.ShouldBindJSON(&body)

		existing.Name = body.Name
		existing.Description = body.Description
		existing.Price = body.Price
		existing.ImageURL = body.ImageURL

		db.Save(&existing)
		c.JSON(200, existing)
	}
}

func DeleteDishHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rid := common.GetRestaurantID(c)
		id := c.Param("id")

		var d Dish
		if db.Where("id = ? AND restaurant_id = ?", id, rid).First(&d).Error != nil {
			c.JSON(404, gin.H{"error": "not found"})
			return
		}

		db.Delete(&d)
		c.JSON(200, gin.H{"message": "deleted"})
	}
}

func SearchDishesHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rid := common.GetRestaurantID(c)
		q := c.Query("q")

		var list []Dish
		db.Where("restaurant_id = ? AND name LIKE ?", rid, "%"+q+"%").
			Find(&list)

		c.JSON(200, list)
	}
}
