package review

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	RestaurantID uint    `json:"restaurant_id" gorm:"index"`
	DishID       uint    `json:"dish_id" gorm:"index"`
	UserID       uint    `json:"user_id"`
	Rating       int     `json:"rating"`
	Comment      string  `json:"comment"`
	Sentiment    float64 `json:"sentiment"`
}
