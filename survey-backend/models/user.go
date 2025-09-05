package models

type User struct {
	ID        int    `json:"id" gorm:"primaryKey;autoIncrement"`
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
}
