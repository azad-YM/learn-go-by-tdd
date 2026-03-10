package romannumeral

import "strings"

func ConvertToArabic(roman string) int {
	total := 0
	for _, row := range romanNumerals {
		for strings.HasPrefix(roman, row.Roman) {
			total += row.Value
			roman = strings.TrimPrefix(roman, row.Roman)
		}
	}

	return total
}
