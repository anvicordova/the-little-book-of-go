## Chapter 1

### How to run a go program?

We have two ways:

1. Run and compile a Go program

```
go run file.go
```

2. Compile a Go program

```
go build file.go
```

Then, run the compiled program

```
./file
```

### What is main?

In Go, the entrypoint of a program in the function main withing a package main

### Imports

We use the `import` keyword to use different libraries. Go requires the libraries imported to be used within the code otherwise it will not compile

```
 package main

 import("fmt)

 func main(){
  fmt.Println("hello")
 }
```

### Variables and declarations

We can declare a variable by using the format `var variableName type`. Depending on the type, go will assign a default value for the variable

```
var age int
```

| Type    | Default value |
| ------- | ------------- |
| int     | 0             |
| string  | ""            |
| boolean | false         |

After the variable declaration we can assign a value to the variable. Note that the type can't be changed.

```
var age int
age = 20
```

Is there are shorter way to do the above? Yes:

We can declare and assign

```
var age int = 20
```

We can declare and assign by using `:=` . In this case, the type will be inferred from the assignation

```
age := 20
```

Other considerations, we can declare two variables of different types at the same time

```
name, age := "Gabi", 20
```

### Function declarations

We can have function with no return values, with one return value or with two returned values

```
func log(message string) {
}

func add(a int, b int) int {
}

func power(name string) (int, bool) {
}
```

When using functions with 2 returned values, we can ignore one of them by using `_`

```
 _, err := myFunction(test)
```

We can reuse `_` for different functions independently of the type returned.
