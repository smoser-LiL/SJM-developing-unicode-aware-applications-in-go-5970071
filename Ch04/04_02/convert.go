package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	in, err := os.Open("holmes.txt")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer in.Close()
	io.CopyN(os.Stdout, in, 100)
}
