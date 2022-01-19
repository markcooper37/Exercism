package allyourbase

import (
    "math"
    "errors"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
        return nil, errors.New("input base must be >= 2")
    }
	if outputBase < 2 {
        return nil, errors.New("output base must be >= 2")
    }
    baseTenNumber := 0
    for i := 0; i < len(inputDigits); i++ {
        if inputDigits[i] < 0 || inputDigits[i] >= inputBase {
            return nil, errors.New("all digits must satisfy 0 <= d < input base")
        }
        baseTenNumber += inputDigits[i] * int(math.Pow(float64(inputBase), float64(len(inputDigits) - i - 1)))
    }
	index := 0
	for i := 1; int(math.Pow(float64(outputBase), float64(i))) <= baseTenNumber; i++ {
        index += 1
    }
	output := []int{}
    for i := index; i >= 0; i-- {
        digit := baseTenNumber / int(math.Pow(float64(outputBase), float64(i)))
        output = append(output, digit)
        baseTenNumber -= digit * int(math.Pow(float64(outputBase), float64(i)))
    }
	return output, nil
}
