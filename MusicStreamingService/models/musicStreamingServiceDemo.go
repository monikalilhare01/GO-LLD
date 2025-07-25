package models

import "fmt"

func Run() {
	service := NewMusicStreamingService()

	user1 := AddUser("1", "John", "abc")
	user2 := AddUser("2", "Doe", "abc")

	service.UserManager.RegisterUser(user1)
	service.UserManager.RegisterUser(user2)

	song1 := &Song{ID: "1", Name: "SONG 1", Artist: "ARTIST 1", Album: "ALBUM 1"}
	song2 := &Song{ID: "2", Name: "SONG 2", Artist: "ARTIST 2", Album: "ALBUM 2"}

	album1 := &Album{ID: "1", Name: "ALBUM 1", Artist: "ARTIST 1", Songs: []*Song{song1}}
	album2 := &Album{ID: "2", Name: "ALBUM 2", Artist: "ARTIST 2", Songs: []*Song{song2}}

	artist1 := &Artist{ID: "1", Name: "Artist 1", Albums: []*Album{album1}}
	artist2 := &Artist{ID: "2", Name: "Artist 2", Albums: []*Album{album2}}

	service.MusicLibrary.AddArtist(artist1)
	service.MusicLibrary.AddArtist(artist2)

	// Login user
	loggedInUser := service.UserManager.LoginUser(user1)
	if loggedInUser != nil {
		fmt.Printf("User logged in: %s\n", loggedInUser.UserName)
	} else {
		fmt.Println("Invalid username or password.")
	}

	// Search songs
	searchResults := service.MusicLibrary.SearchSong("song")
	fmt.Println("\nSearch Results:")
	for _, song := range searchResults {
		fmt.Printf("Song: %s - %s\n", song.Name, song.Artist)

	}
}
