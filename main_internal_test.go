package main

import "testing"

func TestGreet(t *testing.T) {

	type testCase struct {
		lang language
		want string
	}

	var tests = map[string]testCase{

		"English": {
			lang: "en",
			want: "Hello Earth",
		},
		"French": {
			lang: "fr",
			want: "Bonjour le monde",
		},
		"Myanmar": {
			lang: "mm",
			want: "မင်္ဂလာပါ ကမ္ဘာမြေ",
		},
		"Hebrew": {
			lang: "he",
			want: "שלום כדור הארץ",
		},
		"Greek": {
			lang: "el",
			want: "Γεια σου Γη",
		},
		"Latin": {
			lang: "lt",
			want: "Salve Terra",
		},
	}

	for name, tc := range tests {

		t.Run(name, func(t *testing.T) {

			got := greet(tc.lang)

			if got != tc.want {
				t.Errorf("Expected: %q , Got: %q", tc.want, got)
			}
		})
	}
}
