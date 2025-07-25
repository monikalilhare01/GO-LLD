package models

import (
	"strings"
	"sync"
)

type MusicLibrary struct {
	songs  map[string]*Song
	albums map[string]*Album
	artist map[string]*Artist
	mu     sync.RWMutex
}

var (
	libraryInstance *MusicLibrary
	libraryOnce     sync.Once
)

func GetLibraryInstance() *MusicLibrary {
	libraryOnce.Do(func() {
		libraryInstance = &MusicLibrary{
			songs:  make(map[string]*Song),
			albums: make(map[string]*Album),
			artist: make(map[string]*Artist),
		}
	})
	return libraryInstance
}

func (ml *MusicLibrary) AddSong(song *Song) {
	ml.mu.Lock()
	defer ml.mu.Unlock()
	ml.songs[song.ID] = song
}

func (ml *MusicLibrary) AddAlbum(album *Album) {
	ml.mu.Lock()
	defer ml.mu.Unlock()
	ml.albums[album.ID] = album
	for _, song := range album.Songs {
		ml.songs[song.ID] = song
	}
}

func (ml *MusicLibrary) AddArtist(artist *Artist) {
	ml.mu.Lock()
	defer ml.mu.Unlock()
	ml.artist[artist.ID] = artist
	for _, album := range artist.Albums {
		ml.albums[album.ID] = album
		for _, song := range album.Songs {
			ml.songs[song.ID] = song
		}
	}
}

func (ml *MusicLibrary) SearchSong(query string) []*Song {
	ml.mu.Lock()
	defer ml.mu.Unlock()

	query = strings.ToLower(query)
	var results []*Song

	for _, song := range ml.songs {
		if strings.Contains(strings.ToLower(query), song.Name) ||
			strings.Contains(strings.ToLower(query), song.Artist) ||
			strings.Contains(strings.ToLower(query), song.Album) {
			results = append(results, song)
		}
	}
	return results
}
