package main

import "fmt"

func main() {
	// This turns into an error because the slice has length 0
	//scores := make([]int, 0, 10)  slice with length 0, capacity 10
	//scores[7] = 345
	// fmt.Println(scores)

	betterScores := make([]int, 0, 10)
	betterScores = append(betterScores, 5)  // Sets the first element in this case
	betterScores = append(betterScores, 12) // Sets the second element in this case
	fmt.Println(betterScores)

	//Assigning an element in an specifix slice index
	otherScore := make([]int, 0, 10)
	otherScore = otherScore[0:8] // Changes the slice lenght, it goes from 0 to (8-1 == 7)
	otherScore[7] = 1234
	fmt.Println(otherScore) // All the previous indexes have the default, in this case 0
}
