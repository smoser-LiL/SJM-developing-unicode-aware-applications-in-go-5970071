/*
Count how many time each city appers in cities.txt.
You should ignore case.

The right output is:

	asslar: 2
	kraków: 3
	σίπυλοσ: 2
*/
package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/text/cases"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func main() {
	file, err := os.Open("cities.txt")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer file.Close()

	freqs := make(map[string]int) // city -> count

	r := transform.NewReader(file, norm.NFKC)
	fold := cases.Fold()
	s := bufio.NewScanner(r)
	for s.Scan() {
		city := fold.String(s.Text())
		freqs[city] += 1
	}

	if err := s.Err(); err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	for city, count := range freqs {
		fmt.Printf("%s: %d\n", city, count)
	}
}
