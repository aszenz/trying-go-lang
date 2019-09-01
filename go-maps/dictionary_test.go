package maps

import (
	"testing"

	"github.com/aszenz/basic-golang/go-test-helpers"
)

func TestSearch(t *testing.T) {
	dictKey := "test"
	dictValue := "this is just a test"
	dictionary := map[string]string{dictKey: dictValue}

	got := Search(dictionary, dictKey)
	want := dictValue

	testhelpers.AssertStringsEqual(t, got, want)
}
