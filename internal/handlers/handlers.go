package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/gin_gorm_demo/internal/database"
	"github.com/yourusername/gin_gorm_demo/internal/models"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/users", GetUsers)
	r.POST("/users", CreateUser)
}

func GetUsers(c *gin.Context) {
	var users []models.User
	result := database.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户失败"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建用户失败"})
		return
	}

	c.JSON(http.StatusCreated, user)
}
