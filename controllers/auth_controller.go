package controllers

import (
	"smart-command-center-backend/config"
	"smart-command-center-backend/models"
	"smart-command-center-backend/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.BadRequestResponse(c, "Invalid input", err.Error())
		return
	}

	var user models.User
	if err := config.DB.Preload("Role").Where("email = ?", input.Email).First(&user).Error; err != nil {
		utils.UnauthorizedResponse(c, "Invalid email or password")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		utils.UnauthorizedResponse(c, "Invalid email or password")
		return
	}

	token, err := utils.GenerateJWT(user.ID, user.Email, user.Role.Name)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to generate token", err.Error())
		return
	}

	utils.SuccessResponse(c, "Login successful", gin.H{"token": token})
}
