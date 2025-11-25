package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func jwtKey() []byte {
	key := os.Getenv("JWT_SECRET")
	if key == "" {
		key = "default-secret" // Should be changed in production
	}
	return []byte(key)
}

type Claims struct {
	Email  string `json:"email"`
	UserID uint   `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(email string, userID uint) (string, error) {
	exp := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		Email:  email,
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey())
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "missing Authorization header",
			})
			c.Abort()
			return
		}

		// Expect header: "Bearer <token>"
		const prefix = "Bearer "
		if len(authHeader) <= len(prefix) || authHeader[:len(prefix)] != prefix {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid Authorization header format",
			})
			c.Abort()
			return
		}

		tokenStr := authHeader[len(prefix):] // strip "Bearer "

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return jwtKey(), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid or expired token",
			})
			c.Abort()
			return
		}

		// Attach user info
		c.Set("email", claims.Email)
		c.Set("user_id", claims.UserID)

		c.Next()
	}
}
