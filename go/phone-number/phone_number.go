package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
)

func Number(phoneNumber string) (string, error) {
	reg, err := regexp.Compile("[^0-9]")
	if err != nil {
		return "", err
	}
	phoneNumber = reg.ReplaceAllString(phoneNumber, "")
	if len(phoneNumber) != 10 {
		if len(phoneNumber) == 11 && phoneNumber[0] == '1' {
			phoneNumber = phoneNumber[1:]
		} else {
			return "", errors.New("number has incorrect format")
		}
	}
	if phoneNumber[0] == '0' || phoneNumber[3] == '0' || phoneNumber[0] == '1' || phoneNumber[3] == '1' {
		return "", errors.New("number has incorrect format")
	}
	return phoneNumber, nil
}

func AreaCode(phoneNumber string) (string, error) {
	phoneNumber, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return phoneNumber[0:3], nil
}

func Format(phoneNumber string) (string, error) {
	phoneNumber, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", phoneNumber[:3], phoneNumber[3:6], phoneNumber[6:]), nil
}
