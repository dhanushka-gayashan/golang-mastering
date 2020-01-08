package main

func main() {

	nums := []int{1, 2, 3}

	var any interface{}

	any = nums

	//get length
	// any is empty interface type
	// Dynamic value is []int type
	_ = len(any.([]int))
}
