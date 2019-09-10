package maps

import (
	"testing"

	"github.com/aszenz/basic-golang/go-test-helpers"
)

func assertDefinition(t *testing.T, dict Dictionary, word, def string) {
	t.Helper()
	got, ok := dict.Search(word)
	want := def
	testhelpers.AssertNoError(t, ok)
	testhelpers.AssertStringsEqual(t, got, want)
}

func assertNoDefinition(t *testing.T, dict Dictionary, word string, noDefinitionError error) {
	t.Helper()
	r, err := dict.Search(word)
	want := ErrWordDoesNotExist
	testhelpers.AssertError(t, err, want)
	testhelpers.AssertStringsEqual(t, r, "")
}

func TestSearch(t *testing.T) {

	dictKey := "test"
	dictValue := "this is just a test"
	dictionary := Dictionary{dictKey: dictValue}

	t.Run("searching a known word", func(t *testing.T) {
		assertDefinition(t, dictionary, dictKey, dictValue)
	})
	t.Run("searching an unknown word", func(t *testing.T) {
		assertNoDefinition(t, dictionary, "unknown", ErrWordDoesNotExist)
	})

}

func TestAdd(t *testing.T) {
	t.Run("adding a new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "key1"
		definition := "val1"
		dictionary.Add(word, definition)
		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("adding an existing word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "key1"
		definition := "val1"
		anotherDefinition := "val2"
		dictionary.Add(word, definition)
		err := dictionary.Add(word, anotherDefinition)
		testhelpers.AssertError(t, err, ErrAlreadyExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("updating an existing words definition", func(t *testing.T) {
		word := "university"
		oldDefinition := "a place to study"
		dictionary := Dictionary{word: oldDefinition}
		assertDefinition(t, dictionary, word, oldDefinition)

		newDefinition := "a place to learn and study"
		err := dictionary.Update(word, newDefinition)
		assertDefinition(t, dictionary, word, newDefinition)
		testhelpers.AssertNoError(t, err)
	})
	t.Run("updating a non-existent words definition", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Update("castle", "big building")
		testhelpers.AssertError(t, err, ErrTheWordToUpdateDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("deleting an existing word", func(t *testing.T) {
		word := "large"
		definition := "big"
		dictionary := Dictionary{word: definition}
		dictionary.Delete(word)
		assertNoDefinition(t, dictionary, word, ErrWordDoesNotExist)
	})
}
