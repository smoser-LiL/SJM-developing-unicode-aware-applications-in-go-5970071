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
)

func main() {
	file, err := os.Open("cities.txt")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer file.Close()

	freqs := make(map[string]int) // city -> count

	s := bufio.NewScanner(file)
	for s.Scan() {
		freqs[s.Text()] += 1
	}

	if err := s.Err(); err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	for city, count := range freqs {
		fmt.Printf("%s: %d\n", city, count)
	}
}
