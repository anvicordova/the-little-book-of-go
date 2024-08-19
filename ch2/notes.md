# Structures

- Go is not an object-oriented language
- Go support structures, we can define them as follows

```go
type Player struct {
  Name string
  Power int
}
```

## Declarations and Initializations

We can initialize an struct as follows:

```go
player1 := Player{ Name: "Ted", Power: 2000 }
player2 := Player{}
player3 := Player{ "Mir", 3000}
```

In the first way of initialization (player1) we assign values to each of the attributes. With player2 we default to the default values for the types, in this case "" for string and 0 for int. With player3 we omit the name of the fields and rely on the order defined in the structure.

## Passing structs to functions

In Go we pass arguments as values to functions. In the following example, the print statement prints 2000, as the value of power is not changed.

```go
package main

import "fmt"

type Saiyan struct {
	Name  string
	Power int
}

func main() {
	goku := Saiyan{Name: "Goku", Power: 2000}
	incrementPower(goku)
	fmt.Printf("%s's power is %d\n", goku.Name, goku.Power)
}

func incrementPower(s Saiyan) {
	s.Power += 100000
}
```

## Passing a pointer to a function

We can however, pass a pointer of a struct to a function, by declaring the variable as a pointer using `&` and expecting a pointer in the function using `*`. Then the following code will effectively update the power of our struct

```go
package main

import "fmt"

type Saiyan struct {
	Name  string
	Power int
}

func main() {
  // Declaring goku as a pointer
	goku := &Saiyan{Name: "Goku", Power: 2000}
	incrementPower(goku)
	fmt.Printf("%s's power is %d\n", goku.Name, goku.Power)
}

// This function expects a pointer
func incrementPower(s *Saiyan) {
	s.Power += 100000
}
```

## Association of functions to structs

We can easily associate a function to a struct by creating a function that expects the type of the struct as follows.

```go
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

```

As we see in this example the receiver of the function `incrementPower` is a `Player`. Syntaxis:

```go
func (t *StructName) functionName
```

## Different ways to create an struct

We can make use of a constructor to create a new struct as follows:

```golang
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
```

There's no need to return a pointer, we can also return just a value:

```golang
func NewSaiyan(name string, power int) Saiyan {
	return Saiyan{
		Name:  name,
		Power: power,
	}
}
```

We can also make use of the keyword `new``

```golang
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
```

## Struct fields

We can add other structs as fields in a struct. Example:

```golang
type Saiyan struct {
	Name string
	Power int
	Father *Saiyan
}
```

There's no need to have an explicit name for the struct fields, this is also perfectly valid

```golang
type Person struct {
	Name string
}

type Saiyan struct {
	*Person
	Power int
}
```

If we initialize like this we can perfectly do something like

```golang
	goku := &Saiyan{
		Person: &Person{"Goku"},
	}
	fmt.Println(goku.Name)
	fmt.Println(goku.Person.Name)
```
