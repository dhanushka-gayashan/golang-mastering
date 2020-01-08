package main

import "fmt"

func main() {

	names := make([]string, 5)
	fmt.Printf("Length %d Capacity %d Slice %q\n\n", len(names), cap(names), names)

	names = append(names, "einstein", "tesla", "aristo")
	fmt.Printf("Length %d Capacity %d Slice %q\n\n", len(names), cap(names), names)

	names = make([]string, 0, 5)
	names = append(names, "einstein", "tesla", "aristo")
	fmt.Printf("Length %d Capacity %d Slice %q\n\n", len(names), cap(names), names)

	// Even Capacity is 5, only 3 elements available. Last 2 element of names slices have been not "initialized" yet.
	// "Copy" is not initializing elements. Just copy into initialized elements
	// "Append" is initializing new elements and then add the value into the element
	// Therefore, We need to initialize last 2 elements before copy values into those 2 elements
	moreNames := [...]string{"plato", "khayyam", "ptolemy"}

	//copy(names[3:5], moreNames[:2])
	//names = names[:cap(names)]
	names = names[:cap(names)]
	copy(names[3:], moreNames[:2])
	fmt.Printf("Length %d Capacity %d Slice %q\n\n", len(names), cap(names), names)

	clone := make([]string, 3, 5)
	copy(clone, names[len(names)-3:])
	fmt.Printf("Length %d Capacity %d Slice %q\n", len(clone), cap(clone), clone)

	clone = append(clone, names[:2]...)
	fmt.Printf("Length %d Capacity %d Slice %q\n\n", len(clone), cap(clone), clone)

	sliced := clone[1:4:4]
	sliced = append(sliced, "hypatia")
	clone[2] = "elder"
	fmt.Printf("Length %d Capacity %d Slice %q\n", len(clone), cap(clone), clone)
	fmt.Printf("Length %d Capacity %d Slice %q\n\n", len(sliced), cap(sliced), sliced)
}
