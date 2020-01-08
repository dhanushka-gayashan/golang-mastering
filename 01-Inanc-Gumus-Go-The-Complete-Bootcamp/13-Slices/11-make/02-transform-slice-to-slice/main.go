package main

import "fmt"

func main() {

	tasks := []string{"run", "jump", "read"}

	upTasks := make([]string, 0, cap(tasks))
	fmt.Printf("Length %d Capacity %d Tasks %s\n", len(upTasks), cap(upTasks), upTasks)

	for _, t := range tasks {
		upTasks = append(upTasks, t)
		fmt.Printf("Length %d Capacity %d Tasks %s\n", len(upTasks), cap(upTasks), upTasks)
	}
}
