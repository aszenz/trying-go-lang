package testhelpers

import (
	"testing"
	"reflect"
)

// CheckNumbersEqual checks if two float numbers are equal
func CheckNumbersEqual(t *testing.T, got float64, want float64, from string) {
	t.Helper()
	if got != want {
		t.Errorf("expected %.4f but got %.4f from %s", want, got, from)
	}
}

// CheckIntegersEqual checks if two integer numbers are equal
func CheckIntegersEqual(t *testing.T, got int, want int, from string) {
	t.Helper()
	if got != want {
		t.Errorf("expected %d but got %d from %s", want, got, from)
	}
}

// AssertStringsEqual checks if two strings are equal
func AssertStringsEqual(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("expected string \"%s\" but got \"%s\"", want, got)
	}
}

// AssertStringArraysEqual checks if two arrays are equal
func AssertStringArraysEqual(t *testing.T, got []string, want []string) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected array %v but got %v", want, got)
	}
}
