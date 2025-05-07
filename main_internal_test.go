package main

import "testing"

func TestGreet(t *testing.T) {

	want := "Hello Earth"

	got := greet()

	if got != want {

		// test is failed
		t.Errorf("Expected: %q, Got: %q", want, got)
	}
}
