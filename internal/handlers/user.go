package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourorg/doc-agent-demo/internal/models"
)

func ListUsers(c *gin.Context) {
	users := models.GetAllUsers()
	c.JSON(http.StatusOK, gin.H{
		"users": users,
		"count": len(users),
	})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	user, exists := models.GetUserByID(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	// Generate simple ID
	user.ID = fmt.Sprintf("usr_%d", len(models.GetAllUsers())+1)

	models.CreateUser(&user)

	c.JSON(http.StatusCreated, user)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	if !models.UpdateUser(id, &user) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if !models.DeleteUser(id) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted successfully",
	})
}

// GetUserProfile returns detailed user profile including avatar and phone
func GetUserProfile(c *gin.Context) {
	id := c.Param("id")

	user, exists := models.GetUserByID(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	// Return enhanced profile with additional computed fields
	c.JSON(http.StatusOK, gin.H{
		"user": user,
		"profile": gin.H{
			"has_avatar":       user.Avatar != "",
			"has_phone_number": user.PhoneNumber != "",
			"is_admin":         user.Role == "admin",
			"account_age_days": int(user.CreatedAt.Sub(user.CreatedAt).Hours() / 24),
		},
	})
}
