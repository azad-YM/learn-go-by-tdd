package romannumeral

import "strings"

type RomanNumeral struct {
	Value int
	Roman string
}

var romanNumerals = []RomanNumeral{
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {
	var result strings.Builder
	for _, row := range romanNumerals {
		for arabic >= row.Value {
			result.WriteString(row.Roman)
			arabic -= row.Value
		}
	}

	return result.String()
}
