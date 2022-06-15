package repositories

import (
	"momen/entities"

	"gorm.io/gorm"
)


type Repository interface {
	Save(user entities.User)(entities.User, error)
}

type repository struct{
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository  {
	return &repository{db}
}

func (r *repository) Save(user entities.User) (entities.User, error)  {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}