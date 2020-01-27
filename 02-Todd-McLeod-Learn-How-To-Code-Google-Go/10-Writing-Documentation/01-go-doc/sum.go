//Package operations has some math operations
package operations

//Sum method return sum of int array
func Sum(arr ...int) int{
	var total int

	for i := range arr {
		total += i
	}

	return total
}

// go doc
// go doc Sum
// go doc operations.Sum