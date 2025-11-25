package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"orc-shack/internal/user"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/gorm"
)

var googleOAuthConfig = &oauth2.Config{
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
	Scopes:       []string{"email", "profile"},
	Endpoint:     google.Endpoint,
}

func GoogleLoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := googleOAuthConfig.AuthCodeURL("state")
		c.Redirect(http.StatusTemporaryRedirect, url)
	}
}

func GoogleCallbackHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Query("code")

		token, err := googleOAuthConfig.Exchange(context.Background(), code)
		if err != nil {
			c.JSON(400, gin.H{"error": "OAuth exchange failed"})
			return
		}

		// Get profile
		client := googleOAuthConfig.Client(context.Background(), token)
		resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
		if err != nil {
			c.JSON(400, gin.H{"error": "failed to fetch profile"})
			return
		}

		defer resp.Body.Close()

		var profile struct {
			Email string `json:"email"`
			Name  string `json:"name"`
		}
		json.NewDecoder(resp.Body).Decode(&profile)

		// Lookup or create user
		var u user.User
		db.Where("email = ?", profile.Email).First(&u)

		if u.ID == 0 {
			u = user.User{
				Name:         profile.Name,
				Email:        profile.Email,
				PasswordHash: "", // OAuth2 user
			}
			db.Create(&u)
		}

		jwt, _ := GenerateToken(u.Email, u.ID)
		c.JSON(200, gin.H{
			"token": jwt,
		})
	}
}
