package main

import "fmt"

func main() {

	//With Types
	const min int = 1
	const pi float64 = 3.14
	const version string = "2.0.1"
	const debug bool = true
	fmt.Println(min, pi, version, debug)

	// Without Types
	const newMin = 1
	const newPi = 3.14
	const newVersion = "2.0.1"
	const newDebug = true
	fmt.Println(newMin, newPi, newVersion, newDebug)

	// Const Value create with Constant Values
	const constMin = 1 + 1
	const constPi = 3.14 * constMin
	const constVersion = "2.0.1" + "-beta"
	const constDebug = !true
	fmt.Println(constMin, constPi, constVersion, constDebug)
}
