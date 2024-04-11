package main

import (
	"fmt"
)

func main() {
	log("Function with no returned value") // Calling a function with no return value

	result := add(1, 2) // Calling a function with a single return value
	fmt.Printf("Addition result: %d\n", result)

	power, exists := player("goku") // Calling a function with two return values
	if exists == true {
		fmt.Printf("Player has power: %d\n", power)
	}

	_, exists2 := player("error") // Ignoring one of the returned values
	if exists2 == false {
		fmt.Println("Player does not exist")
	}
}

func log(message string) {
	fmt.Printf("Logging: %s\n", message)
}

func add(a, b int) int {
	return a + b
}

func player(name string) (int, bool) {
	if name == "error" {
		return 0, false
	}

	return 3000, true
}
