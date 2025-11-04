package integers

import "testing"

func Test_adder(t *testing.T) {
	t.Run("", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})
}
