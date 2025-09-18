package auth

import (
	"belajar-coding/go/config"
	"belajar-coding/go/model"

	"github.com/google/uuid"
)

func FindUserByEmail(email string) (model.User, error) {
	var user model.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func FindUserByID(id uuid.UUID) (model.User, error) {
	var user model.User
	if err := config.DB.First(&user, "id = ?", id).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}
