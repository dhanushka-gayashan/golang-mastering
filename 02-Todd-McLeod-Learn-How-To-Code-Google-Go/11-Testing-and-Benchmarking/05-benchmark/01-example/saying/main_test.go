package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("Jack")
	if s != "Welcome dear Jack" {
		t.Error("Expected", "Welcome dear Jack", "Got", s)
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Jack"))
	// Output:
	// Welcome dear Jack
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Jack")
	}
}

// go test
// go test -bench .
// go test -bench=.
// go test -bench Greet
// go test -bench=Greet

