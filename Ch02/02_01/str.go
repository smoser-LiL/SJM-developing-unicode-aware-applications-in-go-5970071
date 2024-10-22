package main

import (
	"fmt"
	"unsafe"
)

type String struct {
	str unsafe.Pointer // Underlying byte array
	len int            // length in bytes
}

func main() {
	text := "Go"

	p := unsafe.Pointer(&text)

	s := (*String)(p)
	fmt.Println("len:", s.len)

	c := *(*byte)(s.str)
	fmt.Printf("first char: %c (%d)\n", c, c)
}
