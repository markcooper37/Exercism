package say

import "fmt"

var unitsString = map[int64]string{1: "one", 2: "two", 3: "three", 4: "four", 5: "five", 6: "six", 7: "seven", 8: "eight", 9: "nine"}
var tenToNineteenString = map[int64]string{10: "ten", 11: "eleven", 12: "twelve", 13: "thirteen", 14: "fourteen", 15: "fifteen", 16: "sixteen", 17: "seventeen", 18: "eighteen", 19: "nineteen"}
var tensString = map[int64]string{2: "twenty", 3: "thirty", 4: "forty", 5: "fifty", 6: "sixty", 7: "seventy", 8: "eighty", 9: "ninety"}

func Say(n int64) (string, bool) {
	if n > 999999999999 || n < 0 {
		return "", false
	}
	numberString := ""
	billions := n / 1000000000
	millions := (n % 1000000000) / 1000000
	thousands := (n % 1000000) / 1000
	units := n % 1000
	if billions > 0 {
		numberString = fmt.Sprintf("%s billion", stringSmallNumbers(billions))
		if millions > 0 || thousands > 0 || units > 0 {
			numberString = fmt.Sprintf("%s ", numberString)
		}
	}
	if millions > 0 {
		numberString = fmt.Sprintf("%s%s million", numberString, stringSmallNumbers(millions))
		if thousands > 0 || units > 0 {
			numberString = fmt.Sprintf("%s ", numberString)
		}
	}
	if thousands > 0 {
		numberString = fmt.Sprintf("%s%s thousand", numberString, stringSmallNumbers(thousands))
		if units > 0 {
			numberString = fmt.Sprintf("%s ", numberString)
		}
	}
	if units > 0 {
		numberString = fmt.Sprintf("%s%s", numberString, stringSmallNumbers(units))
	} else if billions == 0 && millions == 0 && thousands == 0 {
		numberString = "zero"
	}
	return numberString, true
}

func stringSmallNumbers(n int64) string {
	numberString := ""
	hundreds := n / 100
	tens := (n % 100) / 10
	units := n % 10

	if hundreds > 0 {
		numberString = fmt.Sprintf("%s hundred", unitsString[hundreds])
		if tens > 0 || units > 0 {
			numberString = fmt.Sprintf("%s ", numberString)
		}
	}
	if tens > 1 {
		numberString = fmt.Sprintf("%s%s", numberString, tensString[tens])
		if units := (n % 10); units >= 1 {
			numberString = fmt.Sprintf("%s-%s", numberString, unitsString[units])
		}
	} else if tens == 1 {
		numberString = fmt.Sprintf("%s%s", numberString, tenToNineteenString[n % 100])
	} else if units != 0 {
		numberString = fmt.Sprintf("%s%s", numberString, unitsString[units])
	}
	return numberString
}
