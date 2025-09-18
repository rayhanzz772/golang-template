package user

import (
	"belajar-coding/go/model"
	"belajar-coding/go/modules/user/dto"
	"belajar-coding/go/utils"
	"context"

	"github.com/google/uuid"
)

func CreateUser(req dto.CreateUserRequest) (model.User, error) {
	user := model.User{
		ID:    uuid.New(),
		Name:  req.Name,
		Email: req.Email,
	}
	if err := CreateUserRepo(&user); err != nil {
		return model.User{}, err
	}
	return user, nil
}

func GetUsers(ctx context.Context, page, perPage int) ([]model.User, utils.Pagination, error) {
	var users []model.User

	offset := (page - 1) * perPage

	total, err := CountUsers(ctx)
	if err != nil {
		return nil, utils.Pagination{}, err
	}

	err = GetUsersRepo(ctx, &users, perPage, offset)
	if err != nil {
		return nil, utils.Pagination{}, err
	}

	meta := utils.NewPagination(page, perPage, total)

	return users, meta, nil
}

func GetUserByID(id uuid.UUID) (model.User, error) {
	var user model.User
	err := GetUserByIDRepo(&user, id)
	return user, err
}

func UpdateUser(user *model.User) error {
	return UpdateUserRepo(user)
}

func DeleteUser(id uuid.UUID) error {
	return DeleteUserRepo(id)
}
