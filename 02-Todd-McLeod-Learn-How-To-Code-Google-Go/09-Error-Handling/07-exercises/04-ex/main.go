package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		lat := "50.2289 N"
		long := "99.4656 W"
		err := fmt.Errorf("more coffee needed - value was %v", f)
		return 0, sqrtError{lat:lat, long:long, err:err}
	}
	return 42, nil
}

