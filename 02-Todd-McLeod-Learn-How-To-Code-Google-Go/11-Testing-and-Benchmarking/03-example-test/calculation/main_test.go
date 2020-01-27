package calculation

import "fmt"

func ExampleSum() {
	fmt.Println(Sum(2, 3))
	// Output:
	// 5
}

func ExampleSumFail() {
	fmt.Println(SumFail(2, 3))
	// Output:
	// 5
}
