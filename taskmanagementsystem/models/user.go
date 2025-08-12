package models

type User struct {
	id    string
	name  string
	email string
}

func NewUser(id string, name, email string) *User {
	return &User{id: id, name: name, email: email}
}

func (u *User) GetID() string {
	return u.id
}

func (u *User) GetName() string {
	return u.name
}
