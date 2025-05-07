package main

import (
	"fmt"
)

func main() {
	greeting := greet("en")
	fmt.Println(greeting)
}

// create a type based on string (type definition)
type language string

func greet(l language) string {

	switch l {
	case "en":
		return "Hello Earth"
	case "fr":
		return "Bonjour le monde"
	default:
		return ""
	}
}
