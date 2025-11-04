package main

import "testing"

func Test_hello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Azad")
		want := "Hello, Azad"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
