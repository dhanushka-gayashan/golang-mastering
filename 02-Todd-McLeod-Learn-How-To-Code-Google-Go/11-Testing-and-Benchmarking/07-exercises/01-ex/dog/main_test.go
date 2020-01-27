package dog

import (
	"fmt"
	"testing"
)

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}

func ExampleYears() {
	fmt.Println(Years(10))
	// Output:
	// 70
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(10))
	// Output:
	// 70
}

func TestYears(t *testing.T) {
	y := Years(10)
	if y != 70 {
		t.Error("Expected", 70, "Got", y)
	}
}

func TestYearsTwo(t *testing.T) {
	y := YearsTwo(10)
	if y != 70 {
		t.Error("Expected", 70, "Got", y)
	}
}

// go test
// go test -bench .
// go test -cover
// go test -coverprofile c.out
// go tool cover -html=c.out
// go -http=:8080