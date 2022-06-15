package services

import (
	"momen/entities"
	"momen/input_post"
	"momen/repositories"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input inputpost.RegisterInput)(entities.User, error)
}

type service struct {
	repository repositories.Repository
}

func NewService(repository repositories.Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input inputpost.RegisterInput)(entities.User, error)  {
	user := entities.User{}
	user.Name = input.Name
	user.Email = input.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(passwordHash)
	newUser, err := s.repository.Save(user)

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}