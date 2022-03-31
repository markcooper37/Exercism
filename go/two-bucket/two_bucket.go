package twobucket

import (
	"errors"
)

type Bucket struct {
	Number   string
	Holding  int
	Capacity int
}

func Solve(sizeBucketOne, sizeBucketTwo, goalAmount int, startBucket string) (string, int, int, error) {
	if sizeBucketOne <= 0 || sizeBucketTwo <= 0 || goalAmount <= 0 {
		return "", 0, 0, errors.New("bucket sizes and goal amount must be greater than 0")
	} else if goalAmount%(HighestCommonFactor(sizeBucketOne, sizeBucketTwo)) != 0 {
		return "", 0, 0, errors.New("goal amount cannot be reached with the given buckets")
	}
	if goalAmount == sizeBucketOne && startBucket == "two" && sizeBucketOne != sizeBucketTwo {
		return "one", 2, sizeBucketTwo, nil
	} else if goalAmount == sizeBucketTwo && startBucket == "one" && sizeBucketOne != sizeBucketTwo {
		return "two", 2, sizeBucketOne, nil
	}
	firstBucket, secondBucket := Bucket{}, Bucket{}
	if startBucket == "one" {
		firstBucket.Number, firstBucket.Capacity, firstBucket.Holding = "one", sizeBucketOne, sizeBucketOne
		secondBucket.Number, secondBucket.Capacity = "two", sizeBucketTwo
		secondBucket.Capacity = sizeBucketTwo
	} else if startBucket == "two" {
		firstBucket.Number, firstBucket.Capacity, firstBucket.Holding = "two", sizeBucketTwo, sizeBucketTwo
		secondBucket.Number, secondBucket.Capacity = "one", sizeBucketOne
	} else {
		return "", 0, 0, errors.New("invalid start bucket")
	}
	actions := 1
	for {
		if firstBucket.Holding == goalAmount {
			return firstBucket.Number, actions, secondBucket.Holding, nil
		} else if secondBucket.Holding == goalAmount {
			return secondBucket.Number, actions, firstBucket.Holding, nil
		}
		if secondBucket.Holding == secondBucket.Capacity {
			secondBucket.Holding = 0
		} else if firstBucket.Holding == 0 {
			firstBucket.Holding = firstBucket.Capacity
		} else {
			var transferAmount int
			if firstBucket.Holding < secondBucket.Capacity-secondBucket.Holding {
				transferAmount = firstBucket.Holding
			} else {
				transferAmount = secondBucket.Capacity - secondBucket.Holding
			}
			firstBucket.Holding -= transferAmount
			secondBucket.Holding += transferAmount
		}
		actions++
	}
}

func HighestCommonFactor(numberOne, numberTwo int) int {
	for numberTwo != 0 {
		numberOne, numberTwo = numberTwo, numberOne%numberTwo
	}
	return numberOne
}
