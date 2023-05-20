package controllers

import (
	"go-crud/initializers"
	"go-crud/models"
	"go-crud/utility"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context)  {
	var body struct {
		FirstName string
		LastName string
		Email string
		Password string
	}

	if err := c.Bind(&body); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	// Password validation
	if !utility.ValidatePassword(body.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid password",
		})
		return
	}

	// Email validation
	if !utility.ValidateEmail(body.Email) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid email",
		})
		return
	}

	// Check if the user already exists
	existingUser := models.User{}
 	result := initializers.DB.Where("email=?", body.Email).First(&existingUser)
	if result.Error == nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "User with provided email already exists",
		})
		return
	}

	// Hashing the users password
	hashedPassword, err := utility.HashingFunction(body.Password)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	user := models.User{
		FirstName: body.FirstName,
		LastName: body.LastName, 
		Email: body.Email, 
		Password: hashedPassword,
	}
	result = initializers.DB.Create(&user)
	
	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func GetUser(c *gin.Context) {
	// Get the req param id
	userID := c.Param("id")

	user := models.User{}
	result := initializers.DB.First(&user, userID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User was not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})

}

func GetAllUsers(c *gin.Context) {
	var users []models.User

	result := initializers.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve users",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})

}

func UpdateUser(c *gin.Context) {
	// Get the req param id
	userID := c.Param("id")

	// Get the existing user from the database
	existingUser := models.User{}
	result := initializers.DB.First(&existingUser, userID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}

	// Bind the request body to a struct
	var updateBody struct {
		FirstName *string `json:"firstName"`
		LastName  *string `json:"lastName"`
		Email     *string `json:"email"`
	}

	if err := c.ShouldBindJSON(&updateBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid payload",
		})
		return
	}

	// Partial update
	if updateBody.FirstName != nil {
		existingUser.FirstName = *updateBody.FirstName
	}
	if updateBody.LastName != nil {
		existingUser.LastName = *updateBody.LastName
	}
	if updateBody.Email != nil {
		existingUser.Email = *updateBody.Email
	}

	// Change the time of update
	now := time.Now()
	existingUser.UpdatedAt = &now

	// Save the updated user to the database
	result = initializers.DB.Save(&existingUser)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": existingUser,
	})
}

func DeleteUser(c *gin.Context) {
	// Get the req param id
	userID := c.Param("id")

	// Check if the user exists
	var existingUser models.User
	result := initializers.DB.First(&existingUser, userID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	// Delete the user
	result = initializers.DB.Delete(&existingUser)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete the user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User was successfully deleted",
	})
}