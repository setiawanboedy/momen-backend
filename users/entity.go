package users

import "time"

type User struct {
	ID             int    `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Name           string `gorm:"size:100;not null"`
	Email          string `gorm:"size:100;not null"`
	PasswordHash   string `gorm:"size:255;not null"`
	AvatarFileName string `gorm:"size:255"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
