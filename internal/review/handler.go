package review

import (
	"orc-shack/internal/common"
	"orc-shack/internal/dish"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateReviewBody struct {
	DishID  uint   `json:"dish_id"`
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}

func CreateReviewHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rid := common.GetRestaurantID(c)
		uid := common.GetUserID(c)

		var body CreateReviewBody
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Ensure dish belongs to restaurant
		var d dish.Dish
		if db.Where("id = ? AND restaurant_id = ?", body.DishID, rid).First(&d).Error != nil {
			c.JSON(404, gin.H{"error": "dish not found"})
			return
		}

		sent := SentimentScore(body.Comment)

		r := Review{
			RestaurantID: uint(rid),
			DishID:       body.DishID,
			UserID:       uid,
			Rating:       body.Rating,
			Comment:      body.Comment,
			Sentiment:    sent,
		}

		db.Create(&r)

		// Recompute avg rating
		var avg float64
		db.Model(&Review{}).
			Where("dish_id = ?", d.ID).
			Select("avg(rating)").Scan(&avg)

		d.AvgRating = avg
		db.Save(&d)

		c.JSON(201, r)
	}
}

func ListReviewsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		did := c.Param("dish_id")

		var list []Review
		db.Where("dish_id = ?", did).Find(&list)

		c.JSON(200, list)
	}
}
