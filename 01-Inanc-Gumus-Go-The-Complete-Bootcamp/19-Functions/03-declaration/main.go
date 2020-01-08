package main

import (
	"fmt"
	"strconv"
)

func main() {

	local := 10

	show(local)

	local = incr(local)
	show(local)

	local = incrBy(local, 5)
	show(local)

	_, err := incrByStr(local, "TWO")
	if err != nil {
		fmt.Printf("err   → %s\n", err)
	}

	local, _ = incrByStr(local, "2")
	show(local)

	show(incrBy(local, 2))
	show(local)

	local = sanitize(incrByStr(local, "NOPE"))
	show(local)

	local = 100
	local = sanitize(incrByStr(local, "2"))
	show(local)

	local = limit(incrBy(local, 5), 2000)
	show(local)
}

func show(local int) {
	fmt.Printf("show -> local = %d\n", local)
}

func incr(local int) int {
	local++
	return local
}

func incrBy(n, factor int) int {
	return n * factor
}

func incrByStr(n int, factor string) (int, error) {
	m, err := strconv.Atoi(factor)
	n = incrBy(n, m)
	return n, err
}

func sanitize(n int, err error) int {

	if err != nil {
		return 0
	}

	return n
}

func limit(n, limit int) (m int) {
	// var m int
	m = n
	if m >= limit {
		m = limit
	}

	// return m
	return
}
