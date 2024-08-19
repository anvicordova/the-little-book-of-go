package main

import "fmt"

func main() {
	scoresArray := [5]int{1, 2, 3, 4, 5}
	slice := scoresArray[2:4]
	slice[0] = 9999
	fmt.Println(scoresArray)
}

// An slice does not create a new array, is a window to the current array
// That is why the original array is modified
