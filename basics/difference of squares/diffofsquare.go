package diffsquares

func SquareOfSum(n int) int {
	Sum := n * (n + 1) / 2
	return Sum * Sum

}

func SumOfSquares(n int) int {
	sumSquares := n * (n + 1) * (2*n + 1) / 6
	return sumSquares
}

func Difference(n int) int {
	Sum := n * (n + 1) / 2
	sumSquares := n * (n + 1) * (2*n + 1) / 6
	return Sum*Sum - sumSquares

}
