package main

import "fmt"

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("here is the error: %v", ce.info)
}
