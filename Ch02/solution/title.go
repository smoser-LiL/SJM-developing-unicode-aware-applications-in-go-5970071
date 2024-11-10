package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

/*
	title returns string with text centered to width and a line below

title("G☺", 6)

	  G☺
	------

Assume width if less than text length.
*/
func title(text string, width int) string {
	size := utf8.RuneCountInString(text)
	padding := (width - size) / 2

	var buf bytes.Buffer
	fmt.Fprintf(&buf, "%s%s\n", strings.Repeat(" ", padding), text)
	fmt.Fprintln(&buf, strings.Repeat("-", width))
	return buf.String()
}

func main() {
	fmt.Println(title("G☺", 6))
}
