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
		song{title: "super sonic", artist: "oasis"},
	}

	rock := playlist{
		genre: "indie rock",
		songs: songs,
	}

	// Clone
	//clone := rock // rock == clone -> Does not work, because songs is Map type.

	fmt.Printf("\n%-20s %20s\n", "TITLE", "ARTIST")
	for _, s := range rock.songs {
		fmt.Printf("%-20s %20s\n", s.title, s.artist)
	}
}
