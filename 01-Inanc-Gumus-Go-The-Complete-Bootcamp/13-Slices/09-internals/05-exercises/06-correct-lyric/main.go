package main

import (
	"fmt"
	"strings"
)

func main() {

	lyric := strings.Fields(`all my troubles seemed so far away oh i believe in yesterday now it looks as though they are here to stay`)

	begin := lyric[11:12]
	lyric = append([]string{begin[0]}, lyric...)

	N, M := 8, 13
	end := lyric[N:M]
	lyric = append(lyric, end...)

	lyric = append(lyric[:N], lyric[M:]...)
	fmt.Printf("%s\n", lyric)
}
