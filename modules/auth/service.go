package auth

import (
	"belajar-coding/go/modules/auth/dto"
	"belajar-coding/go/utils"
)

func ServiceLogin(req dto.LoginRequest) (string, error) {
	user, err := FindUserByEmail(req.Email)
	if err != nil {
		return "", utils.InvalidEmailOrPassword()
	}

	if err := utils.ComparePassword(user.Password, req.Password); err != nil {
		return "", utils.InvalidEmailOrPassword()
	}

	token, err := utils.GenerateToken(user.ID.String(), user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
