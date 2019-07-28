package testhelpers

import "testing"

// CheckNumbersEqual checks if two numbers are equal
func CheckNumbersEqual(t *testing.T, got float64, want float64, from string) {
	t.Helper()
	if got != want {
		t.Errorf("expected %.4f but got %.4f from %s", want, got, from)
	}
}
