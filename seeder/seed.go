package seed

import (
	"belajar-coding/go/config"
	"belajar-coding/go/model"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func SeedUsers() {
	password, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	users := []model.User{
		{
			ID:       uuid.New(),
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: string(password),
		},
	}

	for _, u := range users {
		if err := config.DB.Create(&u).Error; err != nil {
			fmt.Println("Seeder error:", err)
		}
	}
}
