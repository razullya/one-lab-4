package storage

import (
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	User
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		User: newUserStorage(db)}
}
