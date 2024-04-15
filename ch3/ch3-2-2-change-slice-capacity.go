package main

import "fmt"

func main() {
	scores := make([]int, 0, 5) // lenght = 0, capacity = 5
	c := cap(scores)
	fmt.Println(c) // 5

	for i := 0; i < 25; i++ {
		scores = append(scores, i)

		// if our capacity has changed,
		// Go had to grow our array to accommodate the new data
		if cap(scores) != c {
			c = cap(scores)
			fmt.Println(c)
		}
	}
}

// Explanation: With append, Go grows the capacity of the slice by 2X.
// It copies the current values of the slice and creates a new one
// with the double of capacity
