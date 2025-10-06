package user

import (
	"belajar-coding/go/config"
	"belajar-coding/go/model"
	"context"

	"github.com/google/uuid"
)

func CreateUserRepo(user *model.User) error {
	return config.DB.Create(user).Error
}

func GetUserByIDRepo(user *model.User, id uuid.UUID) error {
	return config.DB.First(user, "id = ?", id).Error
}

func UpdateUserRepo(user *model.User) error {
	return config.DB.Updates(user).Error
}

func DeleteUserRepo(id uuid.UUID) error {
	return config.DB.Delete(&model.User{}, id).Error
}

func CountUsers(ctx context.Context) (int, error) {
	var count int64
	if err := config.DB.WithContext(ctx).Model(&model.User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func GetUsersRepo(ctx context.Context, users *[]model.User, limit, offset int) error {
	return config.DB.WithContext(ctx).Limit(limit).Offset(offset).Find(users).Error
}
