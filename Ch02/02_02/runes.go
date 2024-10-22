package main

import (
	"fmt"
)

func main() {
	text := "I ♡ Go"

	fmt.Println("len:", len(text)) // 8 (not 6)
	fmt.Printf("text[0]: %c (%T)\n", text[0], text[0])
	fmt.Printf("text[2]: %c\n", text[2])
}
