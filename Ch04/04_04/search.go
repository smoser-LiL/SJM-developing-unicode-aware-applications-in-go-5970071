package main

import (
	"fmt"
	"time"
)

type Tour struct {
	City string
	Date time.Time
}

// Database
var db []Tour

func insert(city string, year, month, day int) {
	tour := Tour{
		City: city,
		Date: time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC),
	}
	db = append(db, tour)
}

func searchByCity(city string) []Tour {
	var matches []Tour
	for _, t := range db {
		if t.City == city {
			matches = append(matches, t)
		}
	}

	return matches
}

func main() {
	city1 := "Kraków"
	city2 := "Kraków"

	insert(city1, 2025, 5, 17)
	fmt.Println(searchByCity(city2))
}
