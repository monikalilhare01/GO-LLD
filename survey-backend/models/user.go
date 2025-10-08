package models

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserName  string    `json:"user_name"`
	UserEmail string    `json:"user_email"`
	CreatedAt time.Time `json:"created_at"`
}
