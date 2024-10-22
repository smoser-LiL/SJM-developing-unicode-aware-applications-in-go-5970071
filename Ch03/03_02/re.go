package main

import (
	"fmt"
	"regexp"
	"strings"
)

var wordRe = regexp.MustCompile(`[a-zA-Z]+`)

func main() {
	text := "Join us at the city of Kraków for a culinary tour"
	words := wordRe.FindAllString(text, -1)
	fmt.Println(strings.Join(words, "|"))
}
