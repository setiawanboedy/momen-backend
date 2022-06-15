package entities

import "time"

type User struct {
	ID             int
	Name           string
	Email          string
	PasswordHash   string
	AvatarFileName string
	CreatedAt      time.Time
	UpdateAt       time.Time
}
