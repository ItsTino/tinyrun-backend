// internal/api/handlers/runs.go
package handlers

import (
	"tinyrun/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RunHandler struct {
	db *gorm.DB
}

func NewRunHandler(db *gorm.DB) *RunHandler {
	return &RunHandler{db: db}
}

func (h *RunHandler) Create(c *gin.Context) {
	var run models.Run
	if err := c.BindJSON(&run); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	userId, _ := c.Get("user_id")
	run.UserID = userId.(uint)

	if err := h.db.Create(&run).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to save run"})
		return
	}

	c.JSON(200, run)
}

func (h *RunHandler) List(c *gin.Context) {
	userId, _ := c.Get("user_id")
	var runs []models.Run
	h.db.Where("user_id = ?", userId).Find(&runs)
	c.JSON(200, runs)
}
