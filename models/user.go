package models

type User struct {
	Id      string `json:"id,omitempty"`
	Name    string `json:"name"`
	Surname string `json:"surname,omitempty"`
}
