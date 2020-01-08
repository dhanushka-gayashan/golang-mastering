package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	var (
		blender blender
		player  player
		kettle  kettle
		mixer   mixer
	)

	socket := &socket{100}

	fmt.Printf("Socket's available power is %d kW.\n", socket.power)

	if err := socket.plug(blender); err != nil {
		fmt.Println("Blender cannot be powered up:", err)
	}

	if err := socket.plug(player); err != nil {
		fmt.Println("Player cannot be powered up:", err)
	}

	if err := socket.plug(kettle); err != nil {
		fmt.Println("Kettle cannot be powered up:", err)
	}

	if err := socket.plug(mixer); err != nil {
		fmt.Println("Mixer cannot be powered up:", err)
	}

	fmt.Printf("Socket's available power is %d kW.\n", socket.power)
}
