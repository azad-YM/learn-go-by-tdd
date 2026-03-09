package romannumeral

import (
	"fmt"
	"testing"
)

var cases = []struct {
	Arabic int
	Roman  string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{9, "IX"},
	{10, "X"},
	{39, "XXXIX"},
}

func TestRomanNumerals(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gots roman %q", test.Arabic, test.Roman), func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)

			if test.Roman != got {
				t.Errorf("%q got, %q want", got, test.Roman)
			}
		})
	}
}

func TestConvertToArabic(t *testing.T) {
	for _, test := range cases[:3] {
		t.Run(fmt.Sprintf("%q gots arabic %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)

			if test.Arabic != got {
				t.Errorf("%q got, %q want", got, test.Arabic)
			}
		})
	}
}
