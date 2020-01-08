package main

import "fmt"

func main() {

	schools := make([]map[int]string, 2)

	for i := range schools {
		schools[i] = make(map[int]string)
	}

	load(schools[0], []string{"batman", "superman"})
	load(schools[1], []string{"spiderman", "wonder woman"})

	fmt.Println(schools[0])
	fmt.Println(schools[1])
}

// Map always share the Map Header Pointer
// Array Only user for add data. Therefore no need to pass the Pointer
func load(m map[int]string, students []string) {
	for i, name := range students {
		m[i+1] = name
	}
}
