package testhelpers

import "testing"

func AssertError(t *testing.T, gotError error, wantedError error) {
	t.Helper()
	if gotError == nil {
		t.Fatal("wanted an error but didn't get one")
	}
	if gotError != wantedError {
		t.Errorf("wanted %q got %q", gotError, wantedError)
	}
}
func AssertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("wanted no error but got one")
	}
}
