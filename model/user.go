package model

import (
	"go-nextjs-dashboard/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:char(36);not null;unique;primary_key"`
	Name     string    `json:"name" gorm:"type:varchar(255);not null"`
	Email    string    `json:"email" gorm:"type:varchar(255);not null;unique"`
	Password string    `json:"-" gorm:"type:text;not null"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return
	}
	user.Password = hashedPassword
	return
}
