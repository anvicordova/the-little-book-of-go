package main

import "fmt"

func main() {
	scores := make([]int, 5)
	fmt.Printf("capacity =  %d, length =  %d, %v\n", cap(scores), len(scores), scores)

	scores = append(scores, 199)
	fmt.Printf("capacity =  %d, length =  %d, %v\n", cap(scores), len(scores), scores)
}

// Explanation:
// On creation, the slice has capacity and length 5
// The default value for int is 0, therefore the content is already filled
// When using append we are adding a new element to the slice.
// Because of the capacity is full, we will create another slice with the 2X capacity,
// in this case 10.
// The lengh is increased by 1, due the new element
