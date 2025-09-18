package utils

import (
	"belajar-coding/go/config"
	"belajar-coding/go/model"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ComparePassword(hashedPassword, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return err
	}
	return nil
}

func InvalidEmailOrPassword() error {
	return errors.New("invalid email or password")
}

func CheckUserExist(id uuid.UUID) error {
	var user model.User
	return config.DB.First(&user, "id = ?", id).Error
}

func FindUserById(id uuid.UUID) (model.User, error) {
	var user model.User
	if err := config.DB.First(&user, "id = ?", id).Error; err != nil {
		return model.User{}, fmt.Errorf("user not found: %w", err)
	}
	return user, nil
}

func GetUUIDParam(c *fiber.Ctx, param string) (uuid.UUID, error) {
	idStr := c.Params(param)
	return uuid.Parse(idStr)
}
