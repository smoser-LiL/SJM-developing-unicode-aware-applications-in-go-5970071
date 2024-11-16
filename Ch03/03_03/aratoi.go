package numbers

import "fmt"

var arDigits = map[rune]int{
	'٠': 0,
	'١': 1,
	'٢': 2,
	'٣': 3,
	'٤': 4,
	'٥': 5,
	'٦': 6,
	'٧': 7,
	'٨': 8,
	'٩': 9,
}

func arAtoi(s string) (int, error) {
	if s == "" {
		return 0, fmt.Errorf("empty string")
	}

	sign, total := 1, 0

	if s[0] == '-' {
		sign = -1
		s = s[1:]
	}

	for i, c := range s {
		n, ok := arDigits[c]
		if !ok {
			return 0, fmt.Errorf("%q:%d - bad digit", s, i)
		}
		total = total*10 + n
	}

	return total * sign, nil
}
