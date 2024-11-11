package main

import (
	"fmt"
	"regexp"
	"strconv"
)

var numRe = regexp.MustCompile(`\pN+`)

func findNums(text string) ([]int, error) {
	var nums []int
	for _, s := range numRe.FindAllString(text, -1) {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	return nums, nil
}

func main() {
	text := "English: 42, Arabic: ٤٢"
	nums, err := findNums(text)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Println(nums)
}
