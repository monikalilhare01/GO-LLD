package models

import "sync"

type UserManager struct {
	users map[string]*User
	mu    sync.RWMutex
}

func (um *UserManager) RegisterUser(user *User) {
	um.mu.Lock()
	defer um.mu.Unlock()
	um.users[user.ID] = user
}

func (um *UserManager) LoginUser(user *User) *User {
	um.mu.Lock()
	defer um.mu.Unlock()
	for _, u := range um.users {
		if u.UserName == user.UserName && u.Password == user.Password {
			return user
		}
	}
	return nil
}
