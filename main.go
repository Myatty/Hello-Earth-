package main

import (
	"flag"
	"fmt"
)

func main() {

	var lang string

	// Parse the command-line flags and set the value of lang
	flag.StringVar(&lang, "lang", "en", "The required language, e.g. en, ur.")
	flag.Parse()

	greeting := greet(language(lang))
	fmt.Println(greeting)
}

// create a type based on string (type definition)
type language string

// map definition is -> mapName := make(map[keyType]valueType)
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

	/*
		map returns two things
			- the message associated with the key language l - and a
			- boolean (ok as per convention)
	*/
	if !ok {
		return fmt.Sprintf("Unsupported Language: %q", l)
	}

	return greeting
}
