package romannumeral

import "strings"

func ConvertToRoman(arabic int) string {
	conversionTable := []struct {
		Value  int
		Arabic string
	}{
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var result strings.Builder
	for _, row := range conversionTable {
		for arabic >= row.Value {
			result.WriteString(row.Arabic)
			arabic -= row.Value
		}
	}

	return result.String()
}
