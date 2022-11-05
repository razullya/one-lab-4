package service

import (
	"lab4/internal/storage"
	"lab4/models"
)

type User interface {
	CreateUser(name, surname string) error
	GetUserById(id string) (models.User, error)
	UpdateUser(name, surname string, id string) error
	DeleteUser(id string) error
}

type UserService struct {
	stor storage.User
}

func NewUserStorage(stor storage.User) *UserService {
	return &UserService{
		stor: stor,
	}
}
func (u *UserService) CreateUser(name, surname string) error {
	if err := u.stor.CreateUser(name, surname); err != nil {
		return err
	}
	return nil
}

func (u *UserService) GetUserById(id string) (models.User, error) {
	user, err := u.stor.GetUserById(id)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (u *UserService) UpdateUser(name, surname string, id string) error {
	if err := u.stor.UpdateUser(name, surname, id); err != nil {
		return err
	}
	return nil
}
func (u *UserService) DeleteUser(id string) error {
	if err := u.stor.DeleteUser(id); err != nil {
		return err
	}
	return nil
}
