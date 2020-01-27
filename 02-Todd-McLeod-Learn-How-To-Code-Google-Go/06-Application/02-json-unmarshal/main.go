package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last string `json:"Last"`
	Age int `json:"Age"`
}

func main() {

	jData := `
		[
			{
				"First": "James",
				"Last": "Bond",
				"Age": 30
			},
			{
				"First": "Miss",
				"Last": "Moneypenny",
				"Age": 25
			}
		]
	`

	sjData := []byte(jData)
	var people []person

	err := json.Unmarshal(sjData, &people)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(people)
}