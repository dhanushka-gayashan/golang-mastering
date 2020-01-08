package main

import "fmt"

type computer struct {
	brand string
}

func main() {

	var comPtr *computer
	if comPtr == nil {
		fmt.Println("It is nil...")
	}

	comPtr = &computer{brand: "apple"}
	appPtr := comPtr
	if comPtr == appPtr {
		fmt.Printf("Values of First Computer and Apple Computer are equal. FC : %p AC : %p\n", comPtr, appPtr)
	}

	fmt.Println()
	fmt.Println()

	sonyPtr := &computer{brand: "sony"}
	if sonyPtr != appPtr {
		fmt.Printf("Values of Sony Computer and Apple Computer are not equal. FC : %p AC : %p\n", sonyPtr, appPtr)
	}

	fmt.Println()
	fmt.Println()

	apple := *appPtr
	fmt.Printf("Apple Computer Pointer Variable Address : %p\n", &appPtr)
	fmt.Printf("Address of the Apple Computer that Apple Computer Pointer Variable Pointed : %p\n", appPtr)
	fmt.Printf("Address of New Apple Computer Variable : %p\n", &apple)
	if apple == *appPtr {
		fmt.Printf("Values of Apple Computer Pointer and Apple Computer are equal. AP : %v A : %v\n", *appPtr, apple)
	}

	fmt.Println()
	fmt.Println()

	change(sonyPtr, "hp")
	fmt.Printf("Sony Computer New Name : %s\n", sonyPtr.brand)

	fmt.Println()
	fmt.Println()

	newC := valOf(sonyPtr)
	fmt.Printf("New Computer Name is :%v\n", newC.brand)

	fmt.Println()
	fmt.Println()

	fmt.Printf("dell's address            : %p\n", newComputer("dell"))
	fmt.Printf("lenovo's address          : %p\n", newComputer("lenovo"))
	fmt.Printf("acer's address            : %p\n", newComputer("acer"))
}

func change(c *computer, n string) {
	c.brand = n
}

func valOf(c *computer) computer {
	return *c
}

func newComputer(brand string) *computer {
	return &computer{brand: brand}
}
