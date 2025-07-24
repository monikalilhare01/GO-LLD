package models

type Playlist struct {
	ID    string
	Name  string
	Owner *User
	Songs []*Song
}

func CreatePlaylist(id, name string, owner *User) *Playlist {
	return &Playlist{
		ID:    id,
		Name:  name,
		Owner: owner,
		Songs: make([]*Song, 0),
	}
}

func (playlist *Playlist) AddSong(song *Song) {
	playlist.Songs = append(playlist.Songs, song)
}

func (playlist *Playlist) RemoveSong(song *Song) {
	for i, p := range playlist.Songs {
		if p.ID == song.ID {
			playlist.Songs = append(playlist.Songs[:i], playlist.Songs[i+1:]...)
			return
		}
	}
}
