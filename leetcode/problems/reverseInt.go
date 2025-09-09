package problems

import "math"

func reverse(x int) int {

	count := 0

	results := 0
	left := powInt(-2, 31)
	right := powInt(2, 31) - 1

	for x != 0 {

		results = results*10 + x%10
		count++
		x /= 10
		if results < left || results > right {
			return 0
		}
	}

	return results

}
func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
