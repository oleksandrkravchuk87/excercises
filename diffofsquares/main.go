package diffofsquares

func SquareOfSums(n int) int {
	var sum int

	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

func SumOfSquares(n int) int {
	var sumSqur int
	for i := 1; i <= n; i++ {
		sumSqur += i * i
	}
	return sumSqur
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
