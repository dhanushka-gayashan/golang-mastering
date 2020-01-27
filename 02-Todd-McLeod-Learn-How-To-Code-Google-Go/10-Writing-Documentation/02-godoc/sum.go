//Package operations has some math operations
package operations

//Sum function return sum of int array
func Sum(arr ...int) int{
	var total int

	for i := range arr {
		total += i
	}

	return total
}

// godoc -http=:8080