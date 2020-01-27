package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := sqrt(-100)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("norgate math redux:square root of negative number")
		return 0, norgateMathError{"50.50 N", "99.99 N", nme}
	}
	return 42, nil
}
