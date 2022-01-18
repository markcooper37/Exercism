package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
    if n <= 0 {
        return 0, errors.New("Argument must be a positive integer")
    }
    numberOfSteps := 0
	for n != 1 {
        numberOfSteps++
        if n % 2 == 0{
            n = n / 2
        } else {
        	n = 3 * n + 1
        }
    }
	return numberOfSteps, nil
}
