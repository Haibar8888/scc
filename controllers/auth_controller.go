package controllers

import (
	"smart-command-center-backend/services"
	"smart-command-center-backend/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var input services.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.BadRequestResponse(c, "Invalid input", err.Error())
		return
	}

	token, err := services.Login(input)
	if err != nil {
		if err.Error() == "invalid email or password" {
			utils.UnauthorizedResponse(c, err.Error())
		} else {
			utils.InternalServerErrorResponse(c, "Failed to login", err.Error())
		}
		return
	}

	utils.SuccessResponse(c, "Login successful", gin.H{"token": token})
}
