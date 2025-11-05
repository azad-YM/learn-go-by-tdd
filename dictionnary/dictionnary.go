package dictionnary

import "errors"

var (
	ErrNotFound    = errors.New("could not find the word you were looking for")
	ErrWordExisits = errors.New("cannot add word because it already exists")
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	if err == nil {
		return ErrWordExisits
	}

	d[word] = definition
	return nil
}
