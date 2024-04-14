package main

import (
	"fmt"
)

type Saiyan struct {
	Name  string
	Power int
}

func main() {
	player1 := Saiyan{Name: "Run", Power: 2222}
	fmt.Printf("Player 1 name is %s and power is %d\n", player1.Name, player1.Power)

	player2 := Saiyan{}
	fmt.Printf("Player 2 name is %s and power is %d\n", player2.Name, player2.Power)

	player3 := Saiyan{"Kal", 4567}
	fmt.Printf("Player 3 name is %s and power is %d\n", player3.Name, player3.Power)
}
