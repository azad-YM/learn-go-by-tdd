package romannumeral

import (
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	t.Run("1 gots roman I", func(t *testing.T) {
		want := "I"
		got := ConvertToRoman(1)

		if want != got {
			t.Errorf("%q got, %q want", got, want)
		}
	})

	t.Run("2 gots roman II", func(t *testing.T) {
		want := "II"
		got := ConvertToRoman(2)

		if want != got {
			t.Errorf("%q got, %q want", got, want)
		}
	})

	t.Run("3 gots roman III", func(t *testing.T) {
		want := "III"
		got := ConvertToRoman(3)

		if want != got {
			t.Errorf("%q got, %q want", got, want)
		}
	})
}
