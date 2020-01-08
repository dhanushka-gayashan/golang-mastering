package main

import (
	"fmt"
	"strings"
)

func main() {

	slices()
}

func slices() {

	dirs := []string{"up", "down", "left", "right"}
	up(dirs)
	fmt.Printf("slices list   : %p %q\n", &dirs, dirs)

	dirs = []string{"up", "down", "left", "right"}
	upPtr(&dirs)
	fmt.Printf("slices list   : %p %q\n", &dirs, dirs)
}

func up(list []string) {
	// Underlying Data Structure Pointer is already contain in the Slice Header.
	// Therefore no need to send Pointer as the argument
	// Change the Elements in Underlying Array
	for i := range list {
		list[i] = strings.ToUpper(list[i])
		fmt.Printf("up.list[%d]    : %p\n", i, &list[i])
	}

	// Only add new element to local copy
	// Not to the Original Underlying Array
	list = append(list, "HEISEN BUG")
	fmt.Printf("up list      : %p %q\n", &list, list)
}

func upPtr(list *[]string) {
	lv := *list

	for i := range lv {
		lv[i] = strings.ToUpper(lv[i])
	}

	*list = append(*list, "HEISEN BUG")

	fmt.Printf("upPtr list    : %p %q\n", list, list)
}
