package seed

import (
	"belajar-coding/go/config"
	"belajar-coding/go/model"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func SeedUsers() {
	password, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Seeder error:", err)
		return
	}

	users := []model.User{
		{
			ID:       uuid.New(),
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: string(password),
		},
	}

	for _, u := range users {
		result := config.DB.Where("email = ?", u.Email).FirstOrCreate(&u)
		if result.Error != nil {
			fmt.Println("Seeder error:", result.Error)
			continue
		}
		if result.RowsAffected == 0 {
			fmt.Printf("User %s already exists, skipped\n", u.Email)
			continue
		}
		fmt.Printf("User %s created\n", u.Email)
	}
}
