package main

import "fmt"

func main() {
	var scores [3]int
	scores[0] = 20

	fmt.Printf("The size of the scores array is %d\n", len(scores)) // This prints 3

	// Adding a value in [1] instead of [0]

	var grades [5]int
	grades[1] = 30

	fmt.Printf("The size of the scores array is %d\n", len(grades)) // This prints 5
	for index, value := range grades {
		fmt.Printf("Index %d is %d\n", index, value) //Because we are using an int array, Go will add the default value of 0 for each of the elements
	}

	coins := [5]int{10, 20} // Declare and initialize with two values
	for index, value := range coins {
		fmt.Printf("Index %d is %d\n", index, value)
	}

}
