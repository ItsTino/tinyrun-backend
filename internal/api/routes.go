// internal/api/routes.go
package api

import (
	"tinyrun/internal/api/handlers"
	"tinyrun/internal/api/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	authHandler := handlers.NewAuthHandler(db)
	runHandler := handlers.NewRunHandler(db)

	// Auth routes
	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)
	router.GET("/verify", middleware.AuthMiddleware(db), authHandler.Verify)

	// Protected routes
	authorized := router.Group("/")
	authorized.Use(middleware.AuthMiddleware(db))
	{
		authorized.POST("/runs", runHandler.Create)
		authorized.GET("/runs", runHandler.List)
	}

	return router
}
