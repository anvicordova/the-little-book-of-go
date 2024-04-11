package main

import (
	"fmt"
)

func main() {
	var power int
	power = 9000
	fmt.Printf("It's over %d\n", power)

	power2 := getPower2() // declaring and assigning a variable with :=
	fmt.Printf("Second power %d\n", power2)

	name, power3 := "Goku", 23456 // declaring two variables in the same line
	fmt.Printf("%s has this power %d\n", name, power3)
}

func getPower2() int {
	return 45000
}
