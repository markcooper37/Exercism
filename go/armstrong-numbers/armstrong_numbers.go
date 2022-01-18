package armstrong

import (
    "strconv"
    "math"
)

func IsNumber(n int) bool {
	nAsString := strconv.Itoa(n)
    nLength := len(nAsString)
    sum := 0
    for i := 0; i < nLength; i++ {
        digit, _ := strconv.Atoi(string(nAsString[i]))
        sum = sum + int(math.Pow(float64(digit), float64(nLength)))
    }
    return n == sum
}
