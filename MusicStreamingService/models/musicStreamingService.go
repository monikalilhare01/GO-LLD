package models

type MusicStreamingService struct {
	MusicLibrary *MusicLibrary
	UserManager  *UserManager
}

func NewMusicStreamingService() *MusicStreamingService {
	return &MusicStreamingService{
		MusicLibrary: GetLibraryInstance(),
		UserManager:  GetUserManager(),
	}
}
