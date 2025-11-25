package restaurant

import "gorm.io/gorm"

type Restaurant struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}
