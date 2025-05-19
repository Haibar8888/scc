package controllers

import (
	"smart-command-center-backend/config"
	"smart-command-center-backend/models"
	"smart-command-center-backend/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	if err := config.DB.Preload("Role").Find(&users).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to get users", err.Error())
		return
	}
	utils.SuccessResponse(c, "Users retrieved successfully", users)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := config.DB.Preload("Role").First(&user, id).Error; err != nil {
		utils.NotFoundResponse(c, "User not found")
		return
	}
	utils.SuccessResponse(c, "User retrieved successfully", user)
}

func CreateUser(c *gin.Context) {
	type CreateUserInput struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
		RoleID   uint   `json:"role_id" binding:"required"`
	}

	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.BadRequestResponse(c, "Invalid input", err.Error())
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to hash password", err.Error())
		return
	}

	var exists bool
	err = config.DB.Model(&models.Role{}).
		Select("count(*) > 0").
		Where("id = ?", input.RoleID).
		Find(&exists).Error
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to verify role", err.Error())
		return
	}
	if !exists {
		utils.NotFoundResponse(c, "Role not found")
		return
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		RoleID:   input.RoleID,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to create user", err.Error())
		return
	}

	if err := config.DB.Preload("Role").First(&user, user.ID).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to load user role", err.Error())
		return
	}

	utils.CreatedResponse(c, "User created successfully", user)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		utils.NotFoundResponse(c, "User not found")
		return
	}

	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email" binding:"omitempty,email"`
		Password string `json:"password"`
		RoleID   uint   `json:"role_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.BadRequestResponse(c, "Invalid input", err.Error())
		return
	}

	if input.Name != "" {
		user.Name = input.Name
	}
	if input.Email != "" {
		user.Email = input.Email
	}
	if input.RoleID != 0 {
		user.RoleID = input.RoleID
	}
	if input.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			utils.InternalServerErrorResponse(c, "Failed to hash password", err.Error())
			return
		}
		user.Password = string(hashedPassword)
	}

	if err := config.DB.Save(&user).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to update user", err.Error())
		return
	}

	if err := config.DB.Preload("Role").First(&user, user.ID).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to load user role", err.Error())
		return
	}

	utils.SuccessResponse(c, "User updated successfully", user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to delete user", err.Error())
		return
	}
	utils.SuccessResponse(c, "User deleted successfully", nil)
}
