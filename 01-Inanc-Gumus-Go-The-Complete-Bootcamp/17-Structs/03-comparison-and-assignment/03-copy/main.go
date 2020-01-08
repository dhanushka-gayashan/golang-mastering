package main

import "fmt"

type song struct {
	title, artist string
}

type playlist struct {
	genre string
	songs []song
}

func main() {

	songs := []song{
		{title: "wonderwall", artist: "oasis"},
		{title: "super sonic", artist: "oasis"},
	}

	rock := playlist{
		genre: "indie rock",
		songs: songs,
	}

	// Original Value does not change
	// Song variable get a Copy of the original song value
	song := rock.songs[0]
	song.title = "live forever"
	fmt.Println("Rock Song List [0] :", rock.songs[0])
	fmt.Println("Song :", song)

	// Change Original value
	rock.songs[0].title = "live forever"
	fmt.Println("Rock Song List [0] :", rock.songs[0])
}
