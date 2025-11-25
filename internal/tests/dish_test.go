package tests

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"orc-shack/internal/auth"
	"orc-shack/internal/common"
	"orc-shack/internal/dish"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupDishTest() (*gin.Engine, *gorm.DB, string) {
	gin.SetMode(gin.TestMode)

	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&dish.Dish{})

	r := gin.Default()

	token, _ := auth.GenerateToken("test@user.com", 1)

	group := r.Group("/api")
	group.Use(func(c *gin.Context) {
		c.Request.Header.Set("Authorization", token)
		c.Set("email", "test@user.com")
		c.Set("user_id", uint(1))
	})
	group.Use(common.TenantMiddleware())

	group.POST("/dishes", dish.CreateDishHandler(db))
	group.GET("/dishes", dish.ListDishesHandler(db))

	return r, db, token
}

func TestCreateAndListDishes(t *testing.T) {
	r, _, _ := setupDishTest()

	body := `{"name":"Lembas","description":"Elven bread","price":5,"image_url":""}`
	req := bytes.NewBufferString(body)

	w := httptest.NewRecorder()
	req2 := httptest.NewRequest("POST", "/api/dishes", req)
	req2.Header.Set("X-Restaurant-ID", "1")

	r.ServeHTTP(w, req2)

	if w.Code != 201 {
		t.Fatalf("expected 201, got %d", w.Code)
	}
}
