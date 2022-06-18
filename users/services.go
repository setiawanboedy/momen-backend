package users

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterInput)(User, error)
	LoginUser(input LoginInput)(User, error)
	IsEmailAvailable(input CheckEamilInput) (bool, error)
	SaveAvatar(ID int, avatarLocation string) (User, error)
	GetUserByID(ID int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterInput)(User, error)  {
	user := User{}
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

func (s *service) LoginUser(input LoginInput) (User, error){
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash),[]byte(password) )

	if err != nil {
		return user, err
	}

	return user, nil
	
}

func (s *service)IsEmailAvailable(input CheckEamilInput) (bool, error)  {
	email := input.Email

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}
	return true, nil
}

func (s *service) SaveAvatar(ID int, avatarLocation string)(User, error)  {
	user, err := s.repository.FindByID(ID)

	if err != nil {
		return user, err
	}

	user.AvatarFileName = avatarLocation

	userUpdate, err := s.repository.UpdateUser(user)

	if err != nil {
		return userUpdate, err
	}

	return userUpdate, nil
}

func (s *service) GetUserByID(ID int) (User, error)  {
	user, err := s.repository.FindByID(ID)

	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("no user found")
	}

	return user, nil
}