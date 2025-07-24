package models

import "sync"

type User struct {
	ID       string
	UserName string
	Password string
	Playlist []*Playlist
	mu       sync.RWMutex
}

func AddUser(id, userName, password string) *User {
	return &User{
		ID:       id,
		UserName: userName,
		Password: password,
		Playlist: make([]*Playlist, 0),
	}
}

func (u *User) AddPlaylist(playlist *Playlist) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.Playlist = append(u.Playlist, playlist)
}

func (u *User) RemovePlaylist(playlist *Playlist) {
	u.mu.Lock()
	defer u.mu.Unlock()
	for i, p := range u.Playlist {
		if p.ID == playlist.ID {
			u.Playlist = append(u.Playlist[:i], u.Playlist[i+1:]...)
			return
		}
	}
}
