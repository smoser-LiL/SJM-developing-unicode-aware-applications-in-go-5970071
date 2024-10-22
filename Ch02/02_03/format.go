package main

import (
	"fmt"
)

func main() {
	c := 'A'
	fmt.Printf("v: %v\n", c) // 65
	fmt.Printf("c: %c\n", c) // A
	fmt.Printf("c: %v\n", c) // 65

	a, b := 1, "1"
	fmt.Printf("v : a=%v, b=%v\n", a, b)   // a=1, b=1
	fmt.Printf("#v: a=%#v, b=%#v\n", a, b) // a=1, b="1"
}
