package main

import (
	"fmt"
)

func main() {
	greeting := greet("el")
	fmt.Println(greeting)
}

// create a type based on string (type definition)
type language string

var phrasebook = map[language]string{

	"lt": "Salve Terra", // latin
	"el": "Γεια σου Γη", // greek
	"en": "Hello Earth",
	"fr": "Bonjour le monde",   // french
	"he": "שלום כדור הארץ",     // Hebrew
	"mm": "မင်္ဂလာပါ ကမ္ဘာမြေ", // burmese

}

func greet(l language) string {

	greeting, ok := phrasebook[l]

	if !ok {
		return fmt.Sprintf("Unsupported Language: %q", l)
	}

	return greeting
}
