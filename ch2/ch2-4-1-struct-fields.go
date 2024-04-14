package main

import "fmt"

type HappyPlayer struct {
	Name   string
	Power  int
	Father *HappyPlayer
}

func NewHappyPlayer(name string, power int) *HappyPlayer {
	return &HappyPlayer{
		Name:   name,
		Power:  power,
		Father: nil,
	}
}

func (hp *HappyPlayer) AddFather(father *HappyPlayer) {
	hp.Father = father
}

func main() {
	playerWithoutFather := NewHappyPlayer("Run", 5000)

	var fatherName string
	if playerWithoutFather.Father != nil {
		fatherName = playerWithoutFather.Father.Name
	}
	fmt.Printf("%s' power is %d and his father's name is %s\n", playerWithoutFather.Name, playerWithoutFather.Power, fatherName)

	playerWithFather := NewHappyPlayer("Llama", 3000)
	father := NewHappyPlayer("Awesome Father", 10000)
	playerWithFather.AddFather(father)
	fmt.Printf("%s' power is %d and his father's name is %s\n", playerWithFather.Name, playerWithFather.Power, playerWithFather.Father.Name)

}
