package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	roles := []Role{
		{Name: "admin"},
		{Name: "petugas"},
		{Name: "pasien"},
	}

	for _, role := range roles {
		var existing Role
		err := db.Where("name = ?", role.Name).First(&existing).Error
		if err == gorm.ErrRecordNotFound {
			if err := db.Create(&role).Error; err != nil {
				return err
			}
		}
	}

	var adminRole Role
	if err := db.Where("name = ?", "admin").First(&adminRole).Error; err != nil {
		return err
	}

	defaultEmail := "admin@gmail.id"
	var existingUser User
	if err := db.Where("email = ?", defaultEmail).First(&existingUser).Error; err == gorm.ErrRecordNotFound {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		adminUser := User{
			Name:     "Super Admin",
			Email:    defaultEmail,
			Password: string(hashedPassword),
			RoleID:   adminRole.ID,
		}
		if err := db.Create(&adminUser).Error; err != nil {
			return err
		}
	}

	return nil
}
