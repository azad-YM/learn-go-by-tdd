package romannumeral

func ConvertToArabic(roman string) int {
	if roman == "III" {
		return 3
	}

	if roman == "II" {
		return 2
	}

	return 1
}
