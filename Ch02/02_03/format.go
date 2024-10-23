package main

import (
	"fmt"

	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	// Numbers
	n := 1_234.56
	fmt.Println("n:", n)

	en := message.NewPrinter(message.MatchLanguage("en"))
	en.Println("en:", n)
	en.Printf("The cost is $%.2f\n", n)

	// Plural
	const format = "We saw %d birds\n"
	message.Set(language.English, format,
		plural.Selectf(1, "%d",
			"=0", "We saw no birds :(\n",
			plural.One, "We saw a single bird.\n",
			plural.Other, "We saw %d birds.\n",
		),
	)

	printer := message.NewPrinter(language.English)
	printer.Printf(format, 0)
	printer.Printf(format, 1)
	printer.Printf(format, 7)
}
