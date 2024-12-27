// internal/api/handlers/auth.go
package handlers

import (
	"tinyrun/models"
	"tinyrun/pkg/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthHandler struct {
	db *gorm.DB
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{db: db}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword
	user.APIKey = utils.GenerateAPIKey()

	if err := h.db.Create(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "Email already exists"})
		return
	}

	c.JSON(200, gin.H{"api_key": user.APIKey})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	var user models.User
	if err := h.db.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	if !utils.CheckPasswordHash(loginData.Password, user.Password) {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(200, gin.H{"api_key": user.APIKey})
}

func (h *AuthHandler) Verify(c *gin.Context) {
	c.JSON(200, gin.H{"valid": true})
}
