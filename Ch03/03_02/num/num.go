package main

import (
	"fmt"
	"regexp"
	"strings"
)

var numRe = regexp.MustCompile(`\pN+`)

func main() {
	text := "42. The solution is ٤٢"
	nums := numRe.FindAllString(text, -1)
	fmt.Println(strings.Join(nums, "|"))
}
