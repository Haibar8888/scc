package services

import (
	"errors"
	"smart-command-center-backend/config"
	"smart-command-center-backend/models"
	"smart-command-center-backend/utils"

	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Email    string
	Password string
}

func Login(input LoginInput) (string, error) {
	var user models.User
	if err := config.DB.Preload("Role").Where("email = ?", input.Email).First(&user).Error; err != nil {
		return "", errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return "", errors.New("invalid email or password")
	}

	token, err := utils.GenerateJWT(user.ID, user.Email, user.Role.Name)
	if err != nil {
		return "", err
	}

	return token, nil
}
