package main

import "fmt"

type Player struct {
	Name  string
	Power int
}

func (s *Player) incrementPower() {
	s.Power += 1000
}

func main() {
	player1 := Player{Name: "Rin", Power: 300}
	player1.incrementPower()
	fmt.Printf("%s's power is %d\n", player1.Name, player1.Power)
}
