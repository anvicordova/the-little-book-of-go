package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	scores := make([]int, 100)
	for i := 0; i < 100; i++ {
		scores[i] = int(rand.Int31n(1000))
	}

	sort.Ints(scores)
	fmt.Println(scores)

	worst := make([]int, 5)
	copy(worst, scores[:5])
	fmt.Println(worst)
	fmt.Println(scores)

	worst[0] = -1
	fmt.Println(worst)
	fmt.Println(scores)

	copy(worst, scores[10:12]) // It only affects the two first elements of the slice
	fmt.Println(worst)

	// This will affect the original slice, as otherWorst is a window to scores

	otherWorst := scores[:5]
	fmt.Println(otherWorst)
	fmt.Println(scores)

	otherWorst[0] = -1
	fmt.Println(otherWorst)
	fmt.Println(scores)

}

// Explanation: By using copy to fill the slice, the original array is not affected
