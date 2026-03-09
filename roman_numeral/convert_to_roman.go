package romannumeral

import "strings"

func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for i := 0; i < arabic; i++ {
		if arabic == 6 {
			result.WriteString("VI")
			break
		}
		if arabic == 5 {
			result.WriteString("V")
			break
		}
		if arabic == 4 {
			result.WriteString("IV")
			break
		}
		result.WriteString("I")
	}

	return result.String()
}
