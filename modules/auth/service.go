package auth

import (
	"belajar-coding/go/modules/auth/dto"
	"belajar-coding/go/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func serviceLogin(req dto.LoginRequest) (string, error) {
	user, err := FindUserByEmail(req.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		return "", errors.New("invalid email or password")
	}

	token, err := utils.GenerateToken(user.ID.String(), user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
