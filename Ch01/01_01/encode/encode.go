package main

import (
	"fmt"
)

func main() {
	text := "Go"
	for _, c := range text {
		fmt.Printf("%c: %d\n", c, c)
	}
}
