package main

import (
	"fmt"
	"os"

	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
)

func main() {
	data, err := os.ReadFile("holmes.txt")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	_, name, certain := charset.DetermineEncoding(data, "text/plain")
	fmt.Println(name, certain)

	dec := chardet.NewTextDetector()
	result, err := dec.DetectBest(data)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Printf("%s (%d)\n", result.Charset, result.Confidence)
}
