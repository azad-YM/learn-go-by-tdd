package helloworld

import "testing"

func Test_hello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Azad", "")
		want := "Hello, Azad"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Warda", "Spanish")
		want := "Hola, Warda"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("Warda", "French")
		want := "Salut, Warda"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
