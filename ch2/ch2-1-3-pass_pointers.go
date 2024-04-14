package main

import "fmt"

type Saiyan struct {
	Name  string
	Power int
}

func main() {
	goku := &Saiyan{Name: "Goku", Power: 2000}
	incrementPower(goku)
	fmt.Printf("%s's power is %d\n", goku.Name, goku.Power)
}

func incrementPower(s *Saiyan) {
	s.Power += 100000
}
