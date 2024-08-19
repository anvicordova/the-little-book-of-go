package main

import (
	"fmt"
	"strings"
)

func main() {
	haystack := "the spice must flow"
	fmt.Println(strings.Index(haystack[5:], " ")) // From position 5 till the end, find the index of the first " "

	fmt.Println(strings.Index(haystack[:9], "pi")) // From the start until position 9, find the index of the first "pi"
}
