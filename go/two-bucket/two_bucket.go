package twobucket

import "errors"

func Solve(sizeBucketOne, sizeBucketTwo, goalAmount int, startBucket string) (string, int, int, error) {
	if sizeBucketOne <= 0 || sizeBucketTwo <= 0 || goalAmount <= 0 {
		return "", 0, 0, errors.New("bucket sizes and goal amount must be greater than 0")
	}
	return "", 0, 0, nil
}

// Fill the first bucket
// If either bucket has the same size as the required amount, fill if necessary and return
// Otherwise, pour the full bucket into the other bucket since if you emptied it, you would 
// have to refill it, and if you filled the other bucket, you would have to empty it
// Check if either bucket has the required amount, and if so, return
// Otherwise, either fill the first bucket if it is empty or empty the other bucket if it is full
// Pour the first bucket into the other bucket and continue