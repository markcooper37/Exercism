package diffsquares

func SquareOfSum(n int) int {
    // The sum of the first n positive integers is given by n(n+1)/2
	return n*n*(n+1)*(n+1)/4
}

func SumOfSquares(n int) int {
	// The sum of the squares of the first n positive integers is given by n(n+1)(2n+1)/6
    return n*(n+1)*(2*n+1)/6
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
