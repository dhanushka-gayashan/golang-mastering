package main

func main() {

	num := []int{10, 20, 30}

	var many []interface{}

	// many = num won't work. Reason is type mis-match
	// only allow to assign values to single empty interface element

	for _, n := range num {
		many = append(many, n)
	}
}
