package models

type User struct {
	ID        int    `json:"id"`
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
}
