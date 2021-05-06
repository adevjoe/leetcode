// https://leetcode.com/problems/powx-n

package leetcode

func pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	if n == 1 {
		return x
	}
	result := pow(x, n/2)
	if n%2 == 0 {
		result *= result
	} else {
		result = result * result * x
	}
	return result
}

func myPow(x float64, n int) float64 {
	return float64(int(pow(x, n)*100000)) / 100000
}
