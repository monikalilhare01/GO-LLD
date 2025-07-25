package models

import "sync"

type MusicPlayer struct {
	CurrentSong *Song
	IsPlaying   bool
	CurrentTime int
	mu          sync.Mutex
}

func NewMusicPlayer(musicPlayer *MusicPlayer) *MusicPlayer {
	return &MusicPlayer{}
}

func (mp *MusicPlayer) PlaySong(song *Song) {
	mp.mu.Lock()
	defer mp.mu.Unlock()
	mp.CurrentSong = song
	mp.IsPlaying = true
	mp.CurrentTime = 0
}

func (mp *MusicPlayer) PauseSong(song *Song) {
	mp.mu.Lock()
	defer mp.mu.Unlock()
	mp.IsPlaying = false
}

func (mp *MusicPlayer) SeetTo(time int) {
	mp.mu.Lock()
	defer mp.mu.Unlock()
	mp.CurrentTime = time
}
