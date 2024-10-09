package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"yanying-api/models"
)

type Handler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{DB: db}
}

func (h *Handler) Login(c *gin.Context) {
	var req models.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Implement login logic here
	// ...

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func (h *Handler) Register(c *gin.Context) {
	var req models.RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Implement register logic here
	// ...

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

func (h *Handler) CreateCharacter(c *gin.Context) {
	var req models.CreateCharacterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Implement character creation logic here
	// ...

	c.JSON(http.StatusOK, gin.H{"message": "Character created successfully"})
}

func (h *Handler) GetTagList(c *gin.Context) {
	tagType := c.Query("type")

	// Implement tag list retrieval logic here
	// ...

	c.JSON(http.StatusOK, gin.H{"message": "Tag list retrieved successfully"})
}

func (h *Handler) GetWorkerTask(c *gin.Context) {
	// Implement worker task retrieval logic here
	// ...

	c.JSON(http.StatusOK, gin.H{"message": "Worker tasks retrieved successfully"})
}

func (h *Handler) UpdateWorkerTask(c *gin.Context) {
	var req models.UpdateWorkerTaskReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Implement worker task update logic here
	// ...

	c.JSON(http.StatusOK, gin.H{"message": "Worker task updated successfully"})
}