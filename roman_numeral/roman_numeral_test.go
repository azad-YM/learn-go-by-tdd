package romannumeral

import (
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Description string
		Number      int
		Want        string
	}{
		{"1 gots roman I", 1, "I"},
		{"2 gots roman II", 2, "II"},
		{"3 gots roman III", 3, "III"},
		{"4 gots roman IV", 4, "IV"},
		{"5 gots roman V", 5, "V"},
		{"6 gots roman VI", 6, "VI"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Number)

			if test.Want != got {
				t.Errorf("%q got, %q want", got, test.Want)
			}
		})
	}
}
