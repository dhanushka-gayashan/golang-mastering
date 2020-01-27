package word

import (
	"fmt"
	"github.com/dhanukdg/Golang-Udemy-Todd-McLeod-Learn-How-To-Code-Google-Go/11-Testing-and-Benchmarking/07-exercises/02-ex/quote"
	"testing"
)

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

func ExampleCount() {
	fmt.Println(Count("one two three"))
	// Output:
	// 3
}

func TestCount(t *testing.T) {
	n := Count("one two three")
	if n != 3 {
		t.Error("Expected", 3, "Got", n)
	}
}

func TestUseCount(t *testing.T) {
	m := UseCount("one two three three three")
	for k, v := range m {
		switch k {
		case "one":
			if v != 1 {
				t.Error("Expected", 1, "Got", v)
			}
		case "two":
			if v != 1 {
				t.Error("Expected", 1, "Got", v)
			}
		case "three":
			if v != 3 {
				t.Error("Expected", 3, "Got", v)
			}
		}
	}
}

// go test
// go test -bench .
// go test -cover
// go test -coverprofile c.out
// go tool cover -html=c.out
// go -http=:8080