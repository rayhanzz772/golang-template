package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name      string    `gorm:"size:100;not null"`
	Email     string    `gorm:"uniqueIndex;size:150"`
	Password  string    `gorm:"size:255;not null"`
	Status    string    `gorm:"type:enum('active','inactive');default:'active';not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}
	return
}
