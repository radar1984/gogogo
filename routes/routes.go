package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"yanying-api/handlers"
	"yanying-api/middleware"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Inject db to handlers
	h := handlers.NewHandler(db)

	// Public routes
	r.POST("/login", h.Login)
	r.POST("/register", h.Register)

	// Protected routes
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/character/create", h.CreateCharacter)
		auth.GET("/tag/list", h.GetTagList)
		auth.GET("/worker/task", h.GetWorkerTask)
		auth.POST("/worker/task", h.UpdateWorkerTask)
	}

	return r
}