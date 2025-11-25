package auth

import (
	"net/http"
	"orc-shack/internal/common"
	"orc-shack/internal/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegisterBody struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func RegisterHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body RegisterBody
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		hash, _ := HashPassword(body.Password)

		u := user.User{
			Name:         body.Name,
			Email:        body.Email,
			PasswordHash: hash,
		}

		if err := db.Create(&u).Error; err != nil {
			c.JSON(400, gin.H{"error": "email already in use"})
			return
		}

		c.JSON(201, gin.H{"message": "registered"})
	}
}

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body LoginBody
		c.ShouldBindJSON(&body)

		if common.IsBlocked(body.Email) {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "too many failed attempts, try later",
			})
			return
		}

		var u user.User
		if err := db.Where("email = ?", body.Email).First(&u).Error; err != nil {
			common.RegisterFailedLogin(body.Email)
			c.JSON(401, gin.H{"error": "invalid credentials"})
			return
		}

		if !CheckPassword(u.PasswordHash, body.Password) {
			common.RegisterFailedLogin(body.Email)
			c.JSON(401, gin.H{"error": "invalid credentials"})
			return
		}

		common.ResetLoginFailures(body.Email)

		token, _ := GenerateToken(u.Email, u.ID)
		c.JSON(200, gin.H{
			"token": token,
		})
	}
}
