// internal/api/middleware/auth.go
package middleware

import (
	"tinyrun/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("Authorization")
		if apiKey == "" {
			c.JSON(401, gin.H{"error": "No API key provided"})
			c.Abort()
			return
		}

		var user models.User
		if err := db.Where("api_key = ?", apiKey).First(&user).Error; err != nil {
			c.JSON(401, gin.H{"error": "Invalid API key"})
			c.Abort()
			return
		}

		c.Set("user_id", user.ID)
		c.Next()
	}
}
