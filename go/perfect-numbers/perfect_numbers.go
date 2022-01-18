package perfect

import "errors"

type Classification string

const (
	ClassificationDeficient Classification = "Classification Deficient"
    ClassificationPerfect Classification = "Classification Perfect"
    ClassificationAbundant Classification = "Classification Abundant"
)

var ErrOnlyPositive = errors.New("input is not positive")

func Classify(n int64) (Classification, error) {
    if n <= 0 {
        return "", ErrOnlyPositive
    }
    aliquotSum := int64(0)
	for i := int64(1); i < n; i++ {
        if n % i == 0 {
            aliquotSum += i
        }
    }
	if aliquotSum > n {
        return ClassificationAbundant, nil
    } else if aliquotSum == n {
    	return ClassificationPerfect, nil
    } else {
    	return ClassificationDeficient, nil
    }
}
