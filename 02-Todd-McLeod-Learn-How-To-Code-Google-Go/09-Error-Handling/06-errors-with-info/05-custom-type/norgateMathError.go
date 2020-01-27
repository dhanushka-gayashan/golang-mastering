package main

import "fmt"

type norgateMathError struct {
	lat string
	long string
	err error
}

func (n norgateMathError) Error() string {
	return fmt.Sprintf("a norgate math error occured: %v %v %v", n.lat, n.long, n.err)
}