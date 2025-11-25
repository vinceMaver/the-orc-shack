package tests

import (
	"bytes"
	"encoding/json"
	"testing"

	"orc-shack/internal/auth"
	"orc-shack/internal/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupAuthTest() (*gin.Engine, *gorm.DB) {
	gin.SetMode(gin.TestMode)

	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&user.User{})

	r := gin.Default()
	r.POST("/register", auth.RegisterHandler(db))
	r.POST("/login", auth.LoginHandler(db))

	return r, db
}

func TestRegisterAndLogin(t *testing.T) {
	r, _ := setupAuthTest()

	body := `{"name":"Frogo","email":"frogo@shire.me","password":"123456"}`
	w := PerformRequest(r, "POST", "/register", bytes.NewBufferString(body))

	if w.Code != 201 {
		t.Fatalf("expected 201, got %d", w.Code)
	}

	// Login
	w = PerformRequest(r, "POST", "/login", bytes.NewBufferString(body))

	if w.Code != 200 {
		t.Fatalf("login failed, status %d", w.Code)
	}

	var data map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &data)

	if data["token"] == "" {
		t.Fatalf("token missing")
	}
}
