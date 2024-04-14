package main

import "fmt"

type Saiyan struct {
	Name  string
	Power int
}

func main() {
	player := new(Saiyan)
	player.Name = "Lars"
	player.Power = 4500

	fmt.Printf("%s' power is %d\n", player.Name, player.Power)
}
