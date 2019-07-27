package main

import "testing"


func assertEqual(t *testing.T, value1 string, value2 string){
	t.Helper()
	if value1 != value2 {
		t.Errorf("got '%s' want '%s'", value1, value2)
	}
}

func TestHello(test *testing.T) {
	test.Run("says hello to people", func(t *testing.T) {
		got := Hello("aszen", "")
		want := "Hello, aszen"
		assertEqual(t, got, want)
	})
	test.Run("says hello, world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertEqual(t, got, want)
	})
	test.Run("says hello to people, in the nl", func(t *testing.T) {
		got := Hello("aszen", "nl_NL")
		want := "Hallo, aszen"
		assertEqual(t, got, want)
	})
	test.Run("says hello to people, in the fr", func(t *testing.T) {
		got := Hello("aszen", "fr_FR")
		want := "Bonjour, aszen"
		assertEqual(t, got, want)
	})
}
