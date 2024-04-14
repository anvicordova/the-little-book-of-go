package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) FormalName() string {
	return "Dr. " + p.Name
}

type Title struct {
	JobTitle string
}

func (jt *Title) FormalTitle() string {
	return "Title: " + jt.JobTitle
}

type SuperPlayer struct {
	*Person
	JobTitle *Title
	Power    int
}

func NewSuperPlayer(person *Person, power int, jobTitle *Title) *SuperPlayer {
	return &SuperPlayer{
		Person:   person,
		JobTitle: jobTitle,
		Power:    power,
	}
}

func main() {
	person := &Person{Name: "Marita"}
	jobTitle := &Title{JobTitle: "Chef"}
	player := NewSuperPlayer(person, 2000, jobTitle)

	fmt.Printf("%s's power is %d\n", player.Name, player.Power) // We can call player.Name because of composition and the field is not named!
	fmt.Printf("Super Player Formal name is: %s\n", player.FormalName())
	fmt.Printf("Person Formal name is: %s\n", person.FormalName())

	fmt.Printf("Super Player Formal title is: %s\n", player.JobTitle.FormalTitle())
	//fmt.Printf("Super Player Formal title is: %s\n", player.FormalTitle())  This is an error because the field is named!!
}
