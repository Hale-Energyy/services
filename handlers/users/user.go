package handlers

import (
	"fmt"
	"net/http"

	"github.com/axhutoxh/go-starter/config"
	models "github.com/axhutoxh/go-starter/models/users"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func GetUser(c *gin.Context) {
	// Parse user ID from the URL parameter
	id := c.Param("id")

	// Find the user record by ID
	var user models.User
	if err := config.DB.First(&user, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	fmt.Println("Retrieved user: %+v\n", user)

	// Return success response
	c.JSON(http.StatusOK, user)
}

func GetUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func UpdateUser(c *gin.Context) {
	// Get ID from URL param
	id := c.Param("id")

	// Find existing user
	var user models.User
	if err := config.DB.First(&user, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Bind JSON payload to user struct (partial update)
	var input struct {
		FirstName *string  `json:"name"`
		LastName  *string  `json:"last_name"`
		Username  **string `json:"username"`
		Age       *int     `json:"age"`
		Gender    *string  `json:"gender"`
		Height    *float64 `json:"height"`
		Weight    *float64 `json:"weight"`
		BloodType *string  `json:"blood_type"`
		Address   *struct {
			City       *string `json:"city"`
			State      *string `json:"state"`
			Country    *string `json:"country"`
			StreetName *string `json:"street_name"`
			Pincode    *string `json:"pincode"`
		} `json:"address"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields only if provided in input (i.e., partial update)
	if input.FirstName != nil {
		user.FirstName = *input.FirstName
	}

	if input.LastName != nil {
		user.LastName = *input.LastName
	}
	if input.Username != nil {
		user.Username = *input.Username
	}
	if input.Age != nil {
		user.Age = *input.Age
	}
	if input.Gender != nil {
		user.Gender = *input.Gender
	}
	if input.Height != nil {
		user.Height = *input.Height
	}
	if input.Weight != nil {
		user.Weight = *input.Weight
	}
	if input.BloodType != nil {
		user.BloodType = *input.BloodType
	}
	if input.Address != nil {
		if input.Address.City != nil {
			user.Address.City = *input.Address.City
		}
		if input.Address.State != nil {
			user.Address.State = *input.Address.State
		}
		if input.Address.Country != nil {
			user.Address.Country = *input.Address.Country
		}
		if input.Address.StreetName != nil {
			user.Address.StreetName = *input.Address.StreetName
		}
		if input.Address.Pincode != nil {
			user.Address.Pincode = *input.Address.Pincode
		}
	}

	// Save the updated user
	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "user": user})
}

func DeleteUser(c *gin.Context) {
	// Parse user ID from the URL parameter

	id := c.Param("id")
	fmt.Println("Deleting user with ID:", id)
	// Find the user record by ID
	var user models.User
	if err := config.DB.First(&user, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Delete the user record
	if err := config.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
