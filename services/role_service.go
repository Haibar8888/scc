package services

import (
	"errors"
	"smart-command-center-backend/config"
	"smart-command-center-backend/models"
)

func GetAllRoles() ([]models.Role, error) {
	var roles []models.Role
	if err := config.DB.Find(&roles).Error; err != nil {
		return nil, err
	}
	if len(roles) == 0 {
		return nil, errors.New("no roles found")
	}
	return roles, nil
}
