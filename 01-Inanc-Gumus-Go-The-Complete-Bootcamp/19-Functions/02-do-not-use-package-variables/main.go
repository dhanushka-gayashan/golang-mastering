package main

var N int

func main() {

	// other files also can access N variable
	// also, able to change the N variable value from other files
	showN()
	incrN()
	incrN()
	showN()
}

// go run .
