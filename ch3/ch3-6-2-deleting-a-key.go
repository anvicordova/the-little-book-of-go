package main

import "fmt"

func main() {
	myMap := make(map[string]int)
	myMap["scott"] = 345
	myMap["jane"] = 876

	fmt.Println("Map size is:", len(myMap))
	delete(myMap, "jane")
	fmt.Println("Map size afet deletion is:", len(myMap))
}
