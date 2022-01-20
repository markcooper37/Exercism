package romannumerals

import (
    "fmt"
    "errors"
)

func ToRomanNumeral(input int) (string, error) {
    if input > 3000 || input < 1 {
        return "", errors.New("input must be between 1 and 3000")
    }
	units := map[int]string{1: "I", 2: "II", 3: "III", 4: "IV", 5: "V",
                            6: "VI", 7: "VII", 8: "VIII", 9: "IX"}
    tens := map[int]string{1: "X", 2: "XX", 3: "XXX", 4: "XL", 5: "L",
                            6: "LX", 7: "LXX", 8: "LXXX", 9: "XC"}
    hundreds := map[int]string{1: "C", 2: "CC", 3: "CCC", 4: "CD",
                               5: "D", 6: "DC", 7: "DCC", 8: "DCCC", 9: "CM"}
    thousands := map[int]string{1: "M", 2: "MM", 3: "MMM"}
    numeral := ""
    numeral = fmt.Sprintf("%s%s", numeral, thousands[input / 1000])
    numeral = fmt.Sprintf("%s%s", numeral, hundreds[(input % 1000) / 100])
    numeral = fmt.Sprintf("%s%s", numeral, tens[(input % 100) / 10])
    numeral = fmt.Sprintf("%s%s", numeral, units[input % 10])
    return numeral, nil
}
