package main

import "fmt"

func main() {
	lookup := make(map[string]int)
	lookup["goku"] = 9001
	power, exists := lookup["vegeta"]

	fmt.Println(power, exists)
	keysLen := len(lookup)
	fmt.Println("The number of keys is:", keysLen)

}
