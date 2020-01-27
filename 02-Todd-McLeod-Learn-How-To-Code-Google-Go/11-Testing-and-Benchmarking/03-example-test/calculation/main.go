// Package calculation provides functions for basic arithmetic calculations
package calculation

// Sum function give sum of slice of int
func Sum(xi ...int) int {
	var s int
	for _, v := range xi {
		s += v
	}
	return s
}

// SumFail function give wrong sum of slice of int
func SumFail(xi ...int) int {
	var s int
	for _, v := range xi {
		s += v
	}
	return s + 1
}
