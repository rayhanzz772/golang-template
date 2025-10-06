package auth

import (
	"belajar-coding/go/modules/auth/dto"
	"belajar-coding/go/utils"
)

type LoginResponse struct {
	Token string `json:"token"`
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ServiceLogin(req dto.LoginRequest) (*LoginResponse, error) {
	user, err := FindUserByEmail(req.Email)
	if err != nil {
		return nil, utils.InvalidEmailOrPassword()
	}

	if err := utils.ComparePassword(user.Password, req.Password); err != nil {
		return nil, utils.InvalidEmailOrPassword()
	}

	token, err := utils.GenerateToken(user.ID.String(), user.Email)
	if err != nil {
		return nil, err
	}

	response := &LoginResponse{
		Token: token,
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
	}

	return response, nil
}
