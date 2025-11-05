package dictionnary

import "testing"

func Test_dictionnary(t *testing.T) {
	dictionary := Dictionary{"test": "I'm a test!"}

	t.Run("known word", func(t *testing.T) {
		got := dictionary.Search("test")
		want := "I'm a test!"

		assertStrings(t, got, want)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
