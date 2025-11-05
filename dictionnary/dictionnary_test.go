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
