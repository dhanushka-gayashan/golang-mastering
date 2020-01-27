package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {

	pw := "password123"

	hbpw, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
		return
	}

	loginPW := "password1234"
	err = bcrypt.CompareHashAndPassword(hbpw, []byte(loginPW))
	if err != nil{
		fmt.Println("You cannot login....")
	}

	loginPW = "password123"
	err = bcrypt.CompareHashAndPassword(hbpw, []byte(loginPW))
	if err != nil{
		fmt.Println("You cannot login....")
	}else {
		fmt.Println("Login success")
	}
}
