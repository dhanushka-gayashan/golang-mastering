package main

import "fmt"

func main() {

	a, b := 12.5, 25.5

	fmt.Printf("Before Swap Values A %g B %g\n", a, b)
	swapStoredValues(&a, &b)
	fmt.Printf("After Swap Values A %g B %g\n", a, b)

	fmt.Println()
	fmt.Println()

	aPtr, bPtr := &a, &b
	fmt.Printf("Before Swap Addresses A-Ptr %p B-Ptr %p\n", aPtr, bPtr)
	fmt.Printf("Before Swap Values A-Ptr %g B-Ptr %g\n", *aPtr, *bPtr)
	aPtrPtr := &aPtr
	bPtrPtr := &bPtr
	swapStoredAddress(aPtrPtr, bPtrPtr)
	fmt.Printf("After Swap Addresses A-Ptr %p B-Ptr %p\n", aPtr, bPtr)
	fmt.Printf("After Swap Values A-Ptr %g B-Ptr %g\n", *aPtr, *bPtr)

	fmt.Println()
	fmt.Println()

	c, d := 100., 200.
	cPtr, dPtr := &c, &d
	fmt.Printf("Before Swap Addresses C-Ptr %p D-Ptr %p\n", cPtr, dPtr)
	fmt.Printf("Before Swap Values C-Ptr %g D-Ptr %g\n", *cPtr, *dPtr)
	cPtr, dPtr = swapStoredAddressSimple(cPtr, dPtr)
	fmt.Printf("After Swap Addresses C-Ptr %p D-Ptr %p\n", cPtr, dPtr)
	fmt.Printf("After Swap Values C-Ptr %g D-Ptr %g\n", *cPtr, *dPtr)
}

func swapStoredValues(a, b *float64) {
	*a, *b = *b, *a
}

func swapStoredAddress(a, b **float64) {
	**a, **b = **b, **a
}

func swapStoredAddressSimple(c, d *float64) (*float64, *float64) {
	return d, c
}
