package dish

import "gorm.io/gorm"

type Dish struct {
	gorm.Model
	RestaurantID uint    `json:"restaurant_id" gorm:"index"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	ImageURL     string  `json:"image_url"`
	AvgRating    float64 `json:"avg_rating"`
}
