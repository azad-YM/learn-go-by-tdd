package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeat character once", func(t *testing.T) {
		repeated := Repeat("a", 1)
		expected := "a"

		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("repeat character twice", func(t *testing.T) {
		got := Repeat("a", 2)
		expected := "aa"

		assertCorrectMessage(t, got, expected)
	})
}

func assertCorrectMessage(t testing.TB, got, expected string) {
	t.Helper()
	if got != expected {
		t.Errorf("expected %q but got %q", expected, got)
	}
}
