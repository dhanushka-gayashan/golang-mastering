package main

import (
	"encoding/json"
	"fmt"
)

type permission map[string]bool

type user struct {
	Name       string     `json:"username"`
	Password   string     `json:"-"`
	Permission permission `json:"perms,omitempty"`
}

func main() {

	users := []user{
		{"nick", "1234", nil},
		{"jack", "5673", permission{"admin": true}},
		{"ben", "7893", permission{"write": true}},
	}

	// Encoding to Json - Marshal
	out, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(out))
}

// Save output into a file
// go run main.go > users.json
