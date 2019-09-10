package maps

const (
	// ErrWordDoesNotExist is thrown while searching for a word which doesn't exist
	ErrWordDoesNotExist = DictionaryError("cannot find this words definition")

	// ErrAlreadyExists is thrown when adding a prexisting word
	ErrAlreadyExists = DictionaryError("word already exists")

	// ErrTheWordToUpdateDoesNotExist is thrown when performing an update on a non existent word
	ErrTheWordToUpdateDoesNotExist = DictionaryError("The word you are trying to update does not exist")
)

// Dictionary stores string key value pairs
type Dictionary map[string]string

// DictionaryError are Dictionary Errors
type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

// Search searches for a given key
func (d Dictionary) Search(key string) (string, error) {
	v, ok := d[key]
	if !ok {
		return "", ErrWordDoesNotExist
	}
	return v, nil
}

// Add adds a new key value pair to the dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrWordDoesNotExist:
		d[word] = definition
	case nil:
		return ErrAlreadyExists
	default:
		return err
	}
	return nil
}

// Update updates the definition of an existing word
func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrWordDoesNotExist:
		return ErrTheWordToUpdateDoesNotExist
	case nil:
		d[word] = newDefinition
	default:
		return err
	}
	return nil
}

// Delete deletes the word from the dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
