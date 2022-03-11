package cipher

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

type shift int
type vigenere string

func NewCaesar() Cipher {
	return shift(3)
}

func NewShift(distance int) Cipher {
	if distance == 0 || distance < -25 || distance > 25 {
		return nil
	}
	return shift(distance)
}

func (c shift) Encode(input string) string {
	reg, err := regexp.Compile("[^a-zA-Z]")
	if err != nil {
		log.Fatal(err)
	}
	input = reg.ReplaceAllString(input, "")
	input = strings.ToLower(input)
	output := ""
	for _, char := range input {
		char += int32(c)
		if char < 97 {
			char += 26
		} else if char > 122 {
			char -= 26
		}
		output = fmt.Sprintf("%s%s", output, string(char))
	}
	return output
}

func (c shift) Decode(input string) string {
	output := ""
	for _, char := range input {
		char -= int32(c)
		if char < 97 {
			char += 26
		} else if char > 122 {
			char -= 26
		}
		output = fmt.Sprintf("%s%s", output, string(char))
	}
	return output
}

func NewVigenere(key string) Cipher {
	reg, err := regexp.Compile("[^a-z]")
	if err != nil {
		log.Fatal(err)
	}
	onlyLowerCaseKey := reg.ReplaceAllString(key, "")
	if len(onlyLowerCaseKey) != len(key) {
		return nil
	}
	for _, char := range key {
		if char != 'a' {
			return vigenere(key)
		}
	} 
	return nil
}

func (v vigenere) Encode(input string) string {
	reg, err := regexp.Compile("[^a-zA-Z]")
	if err != nil {
		log.Fatal(err)
	}
	input = reg.ReplaceAllString(input, "")
	input = strings.ToLower(input)
	output := ""
	for index, char := range input {
		char += int32(v[index % len(v)] - 97)
		if char < 97 {
			char += 26
		} else if char > 122 {
			char -= 26
		}
		output = fmt.Sprintf("%s%s", output, string(char))
	}
	return output
}

func (v vigenere) Decode(input string) string {
	output := ""
	for index, char := range input {
		char -= int32(v[index % len(v)] - 97)
		if char < 97 {
			char += 26
		} else if char > 122 {
			char -= 26
		}
		output = fmt.Sprintf("%s%s", output, string(char))
	}
	return output
}
