package controllers

import (
	"smart-command-center-backend/services"
	"smart-command-center-backend/utils"

	"github.com/gin-gonic/gin"
)

func GetAllRoles(c *gin.Context) {
	roles, err := services.GetAllRoles()
	if err != nil {
		if err.Error() == "no roles found" {
			utils.NotFoundResponse(c, err.Error())
		} else {
			utils.InternalServerErrorResponse(c, "Failed to retrieve roles", err.Error())
		}
		return
	}
	utils.SuccessResponse(c, "Roles retrieved successfully", roles)
}
