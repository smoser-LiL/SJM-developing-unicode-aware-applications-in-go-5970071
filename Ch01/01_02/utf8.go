package main

import (
	"fmt"
)

func main() {
	text := "I ♡ Go"
	for _, c := range text {
		fmt.Printf("%c: 0x%x\n", c, c)
	}
}
