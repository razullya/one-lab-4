package service

import "lab4/internal/storage"

type Service struct {
	User
}

func NewService(stor *storage.Storage) *Service {
	return &Service{
		User: NewUserStorage(stor.User),
	}
}
