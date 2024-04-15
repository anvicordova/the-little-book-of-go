# Arrays, Slices and Maps

## Arrays

Arrays in Go are not dynamic. We need to declare in advance the size of it.

```golang
var scores [10]int
scores[0] = 500
```

We can also initialize with values as follows

```golang
scores := [5]int{10,20,30,40}
```

We can have the size of an array by using `len`

```golang
len(scores)
```

We can iterate over an array by using `range`

```golang
for index, value := range scores{
fmt.Printf("Index %d is %d\n", index, value)
}
```

## Slices

Slices are structures that represent portions of an array. We can declare them in different ways:

```golang
scores := []int{1,2,3,4,5}
```

or

```golang
scores := make([]int, 10) // Initializes the slice with a length of 10 and capacity of 10
```

or

```golang
scores := make([]int, 0, 10) // Initializes the slice with a length of 0 and capacity of 10
```

or

```golang
var names []string // Nil slice, should be used with append. Use it when the number of elements is unknown
```

What is the difference between length and capacity?

Length --> Refers to the size of the slice

```golang
len(slice) // Returns the length of a slice
```

Capacity --> Referes to the size of the underlying array

```golang
cap(slice) // Returns the capacity of a slice
```
