package main

import "fmt"

type Saiyan struct {
	Name  string
	Power int
}

func NewSaiyan(name string, power int) *Saiyan {
	return &Saiyan{
		Name:  name,
		Power: power,
	}
}

func main() {
	player := NewSaiyan("Rumi", 4500)
	fmt.Printf("%s' power is %d\n", player.Name, player.Power)
}
