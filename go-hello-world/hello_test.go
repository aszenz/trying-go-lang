package main

import (
	"testing"

	"github.com/aszenz/basic-golang/go-test-helpers"
)

func TestHello(test *testing.T) {
	test.Run("says hello to people", func(t *testing.T) {
		got := Hello("aszen", "")
		want := "Hello, aszen"
		testhelpers.AssertStringsEqual(t, got, want)
	})
	test.Run("says hello, world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		testhelpers.AssertStringsEqual(t, got, want)
	})
	test.Run("says hello to people, in the nl", func(t *testing.T) {
		got := Hello("aszen", "nl_NL")
		want := "Hallo, aszen"
		testhelpers.AssertStringsEqual(t, got, want)
	})
	test.Run("says hello to people, in the fr", func(t *testing.T) {
		got := Hello("aszen", "fr_FR")
		want := "Bonjour, aszen"
		testhelpers.AssertStringsEqual(t, got, want)
	})
}
