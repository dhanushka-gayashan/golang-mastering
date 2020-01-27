package main

import (
	"testing"
)

func TestMySum(t *testing.T) {
	x := mySum(2, 3)
	if x != 5 {
		t.Error("Expected", 5, "Got", x)
	}
}

func TestMySumFail(t *testing.T) {
	x := mySumFail(2, 3)
	if x != 5 {
		t.Error("Expected", 5, "Got", x)
	}
}

// go test
// go test -v
