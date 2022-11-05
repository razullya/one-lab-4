package storage

import (
	"fmt"
	"lab4/models"

	"github.com/jmoiron/sqlx"
)

type User interface {
	CreateUser(name, surname string) error
	GetUserById(id string) (models.User, error)
	UpdateUser(name, surname string, id string) error
	DeleteUser(id string) error
}
type UserStorage struct {
	db *sqlx.DB
}

func newUserStorage(db *sqlx.DB) *UserStorage {
	return &UserStorage{
		db: db,
	}
}

func (u *UserStorage) CreateUser(name, surname string) error {
	query := `INSERT INTO user(name, surname) VALUES ($1, $2);`
	_, err := u.db.Exec(query, name, surname)
	if err != nil {
		return fmt.Errorf("storage: create user: %w", err)
	}
	return nil
}

func (u *UserStorage) GetUserById(id string) (models.User, error) {
	query := `SELECT name, surname FROM user WHERE id=$1;`
	row := u.db.QueryRow(query, id)
	var user models.User
	err := row.Scan(&user.Name, &user.Surname)
	if err != nil {
		return models.User{}, fmt.Errorf("storage: get user by login: %w", err)
	}
	return user, nil
}

func (u *UserStorage) UpdateUser(name, surname string, id string) error {
	query := `UPDATE user SET name = $1, surname = $2 WHERE id = $3;`
	_, err := u.db.Exec(query, name, surname, id)
	if err != nil {
		return fmt.Errorf("storage: delete post: %w", err)
	}
	return nil
}

func (u *UserStorage) DeleteUser(id string) error {
	query := `DELETE FROM user WHERE id = $1;`
	_, err := u.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("storage: delete post: %w", err)
	}
	return nil
}
