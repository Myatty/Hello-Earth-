package main

import "testing"

func TestGreet_English(t *testing.T) {

	lang := language("en")
	want := "Hello Earth"

	got := greet(lang)

	if got != want {
		// test is failed
		t.Errorf("Expected: %q, Got: %q", want, got)
	}
}

func TestGreet_French(t *testing.T) {

	lang := language("fr")
	want := "Bonjour le monde"

	got := greet(lang)

	if got != want {
		t.Errorf("Expected: %q, Got: %q", want, got)
	}
}

func TestGreet_Burmese(t *testing.T) {

	lang := language("mm")
	want := "မင်္ဂလာပါ ကမ္ဘာမြေ"

	got := greet(lang)

	if got != want {
		t.Errorf("Expected: %q, Got: %q", want, got)
	}
}
