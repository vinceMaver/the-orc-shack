package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"orc-shack/internal/auth"
	"orc-shack/internal/common"
	"orc-shack/internal/dish"
	"orc-shack/internal/restaurant"
	"orc-shack/internal/review"
	"orc-shack/internal/user"
)

func migrate(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&restaurant.Restaurant{})
	db.AutoMigrate(&dish.Dish{})
	db.AutoMigrate(&review.Review{})
	log.Println("Migrations completed.")
}

func main() {
	// --- Init ---
	cfg := common.LoadConfig()
	db := common.InitDB()
	migrate(db)

	common.InitCache()

	// --- Router ---
	r := gin.Default()
	r.Use(common.LoggingMiddleware())

	// --- Public routes ---
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", auth.RegisterHandler(db))
		authGroup.POST("/login", auth.LoginHandler(db))

		// Optional OAuth2
		authGroup.GET("/google", auth.GoogleLoginHandler())
		authGroup.GET("/google/callback", auth.GoogleCallbackHandler(db))
	}

	// --- Protected routes (JWT required) ---
	api := r.Group("/api")
	api.Use(auth.JWTAuthMiddleware())

	// Rate limit all authenticated routes
	api.Use(common.RateLimitMiddleware(60)) // 60 req/min per user

	// Multi-tenant — all routes require restaurant ID
	api.Use(common.TenantMiddleware())

	// Restaurant management
	rest := api.Group("/restaurants")
	{
		rest.POST("/", restaurant.CreateRestaurantHandler(db))
		rest.GET("/", restaurant.ListRestaurantsHandler(db))
	}

	// Dishes
	dishes := api.Group("/dishes")
	{
		dishes.POST("/", dish.CreateDishHandler(db))
		dishes.GET("/", dish.ListDishesHandler(db))
		dishes.GET("/search", dish.SearchDishesHandler(db))
		dishes.GET("/:id", dish.GetDishHandler(db))
		dishes.PUT("/:id", dish.UpdateDishHandler(db))
		dishes.DELETE("/:id", dish.DeleteDishHandler(db))
	}

	// Reviews
	reviews := api.Group("/reviews")
	{
		reviews.POST("/", review.CreateReviewHandler(db))
		reviews.GET("/dish/:dish_id", review.ListReviewsHandler(db))
	}

	r.Run(":" + cfg.Port)
}
