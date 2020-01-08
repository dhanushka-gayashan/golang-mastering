package main

import "fmt"

func main() {

	proBooks := [5]string{"Java", "Scala", "Golang", "Python", "Kotlin"}
	webBooks := [5]string{"HTML", "CSS", "JS", "JWT", "Less"}

	fmt.Printf("proBooks			: %#v\n", proBooks)
	fmt.Printf("proBooks			: %T\n", proBooks)
	fmt.Printf("proBooks Length		: %d\n", len(proBooks))

	fmt.Println()

	fmt.Printf("webBooks			: %#v\n", webBooks)
	fmt.Printf("webBooks			: %T\n", webBooks)
	fmt.Printf("webBooks Length		: %d\n", len(webBooks))

	fmt.Println()

	if proBooks != webBooks {
		fmt.Println("proBooks == webBooks", proBooks == webBooks)
		fmt.Println("proBooks can compare with webBooks, because both are [5]string type")
	}

	fmt.Println()

	proCourses := []string{"Java", "Scala", "Golang", "Python", "Kotlin"}
	webCourses := []string{"HTML", "CSS", "JS", "JWT", "Less"}

	fmt.Printf("proCourses			: %#v\n", proCourses)
	fmt.Printf("proCourses			: %T\n", proCourses)
	fmt.Printf("proCourses Length	: %d\n", len(proCourses))

	fmt.Println()

	fmt.Printf("webCourses			: %#v\n", webCourses)
	fmt.Printf("webCourses			: %T\n", webCourses)
	fmt.Printf("webCourses Length	: %d\n", len(webCourses))

	fmt.Println()

	if proCourses != nil && webCourses != nil {
		fmt.Println("proCourses == nil", proCourses == nil)
		fmt.Println("webCourses == nil", webCourses == nil)
		fmt.Println("Slice can only compare with nil.")
		fmt.Println("Cannot compare proCourses == webCourses")
	}

	fmt.Println()

	// Compare 2 Slices
	//=================
	var ok string

	if len(proCourses) != len(webCourses) {
		ok = "not "
	} else {
		for i, pc := range proCourses {
			if pc != webCourses[i] {
				ok = "not "
				break
			}
		}
	}

	fmt.Printf("proCourse and webCourse are %sequal\n\n", ok)
}
