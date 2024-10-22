package main

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/text/search"
)

func main() {
	g1, g2 := "Go", "go"

	fmt.Println("ToLower: Go == go:", strings.ToLower(g1) == strings.ToLower(g2)) // true

	s1, s2, s3 := "Σ", "σ", "ς"

	l1, l2 := strings.ToLower(s1), strings.ToLower(s2)
	fmt.Printf("ToLower: %s == %s %v\n", s1, s2, l1 == l2) // true
	l3 := strings.ToLower(s3)
	fmt.Printf("ToLower: %s == %s %v\n", s1, s3, l1 == l3) // false

	fmt.Printf("EqualFold: %s == %s %v\n", s1, s2, strings.EqualFold(s1, s2))
	fmt.Printf("EqualFold: %s == %s %v\n", s1, s3, strings.EqualFold(s1, s3))

	fold := cases.Fold()
	for _, s := range []string{"Σ", "σ", "ς"} {
		fmt.Printf("fold: %s -> %s\n", s, fold.String(s))
	}

	r1 := "der Fluß" // ß = Eszett
	r2 := "der Fluss"
	fmt.Println("EqualFold: ß == ss", strings.EqualFold(r1, r2))

	matcher := search.New(language.German, search.Loose)
	fmt.Println("matcher:   ß == ss", matcher.EqualString(r1, r2))
}
