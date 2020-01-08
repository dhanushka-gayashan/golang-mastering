package main

type computer struct {
	brand string
}

// Count Variables
// a and b in main()
// c twice because change() call 2 times
func main() {
	a := &computer{"Apple"}
	b := a
	change(b)
	change(b)
}

func change(c *computer) {
	c.brand = "Indie"
	c = nil
}
