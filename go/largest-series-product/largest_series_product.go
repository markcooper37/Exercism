package lsproduct

import (
    "strconv"
    "errors"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
    for i := 0; i < len(digits); i++ {
        if int(digits[i]) < 48 || int(digits[i]) > 57 {
            return -1, errors.New("digits input must only contain digits")
        }
    }

    if span < 0 {
        return -1, errors.New("span must not be negative")
    }

    if span == 0 {
        return 1, nil
    }

	if span > len(digits) {
        return -1, errors.New("span must be smaller than string length")
    }

	iterations := len(digits) - span + 1
    var start int64 = 0
    for i := 0; i < iterations; i++ {
        var product int64 = 1
        for j := i; j < i + span; j++ {
            digit, _ := strconv.Atoi(string(digits[j]))
            product = product * int64(digit)
        }
    	if product > start {
        	start = product
        }
    }
	return start, nil
}
