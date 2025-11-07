package dictionnary

import "testing"

func Test_Search(t *testing.T) {
	dictionary := Dictionary{"test": "I'm a test!"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "I'm a test!"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrNotFound.Error()

		assertError(t, err)
		assertStrings(t, err.Error(), want)
	})
}

func Test_Add(t *testing.T) {
	dictionary := Dictionary{"existing": "I'm a existing word"}

	t.Run("new word", func(t *testing.T) {
		word := "new"
		definition := "I'm a new word in dictionnary!"

		dictionary.Add(word, definition)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "existing"
		definition := "I'm a existing word"

		err := dictionary.Add(word, definition)
		assertError(t, err)
		assertStrings(t, err.Error(), ErrWordExisits.Error())
	})
}

// modifier la définition d'un mot
// supprimer un mot

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, got, definition)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("expected to get an error.")
	}
}
