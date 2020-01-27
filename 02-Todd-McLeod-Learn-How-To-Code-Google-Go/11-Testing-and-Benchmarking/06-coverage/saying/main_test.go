package saying

import (
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("Jack")
	if s != "Hello Jack" {
		t.Error("Expected", "Hello Jack", "Got", s)
	}
}

// go test -cover
// go test -coverprofile c.out
// go tool cover -html=c.out
