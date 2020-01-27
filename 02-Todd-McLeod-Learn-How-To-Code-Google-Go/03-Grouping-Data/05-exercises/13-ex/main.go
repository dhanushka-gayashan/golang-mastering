package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	t := truck{
		vehicle:   vehicle{2, "black"},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{4, "red"},
		luxury:  false,
	}

	fmt.Println("Truck", t)
	fmt.Println("Sedan", s)

	fmt.Println("Truck have ", t.doors, " doors")
	fmt.Println("Sedan have ", s.doors, " doors")
}
