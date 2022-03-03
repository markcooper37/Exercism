package cryptosquare

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strings"
)

func Encode(pt string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	pt = reg.ReplaceAllString(pt, "")
	pt = strings.ToLower(pt)
	c := int(math.Ceil(math.Sqrt(float64(len(pt)))))
	r := int(math.Ceil(float64(len(pt))/float64(c)))
	code := ""
	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			pos := i + j*c
			if pos >= len(pt) {
				code = fmt.Sprintf("%s%s", code, " ")
			} else {
				code = fmt.Sprintf("%s%s", code, string(pt[pos]))
			}
		}
		if i < c-1 {
			code = fmt.Sprintf("%s%s", code, " ")
		}
	}
	return code
}
