package controllers

import (
	"smart-command-center-backend/config"
	"smart-command-center-backend/models"
	"smart-command-center-backend/utils"

	"github.com/gin-gonic/gin"
)

func GetAllRoles(c *gin.Context) {
	var roles []models.Role

	if err := config.DB.Find(&roles).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to retrieve roles", err)
		return
	}

	if len(roles) == 0 {
		utils.NotFoundResponse(c, "No roles found")
		return
	}

	utils.SuccessResponse(c, "Roles retrieved successfully", roles)
}
