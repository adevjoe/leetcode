// https://leetcode.com/problems/ugly-number-ii/

package algorithms

func nthUglyNumber(n int) int {
	if n == 1 {
		return n
	}
	got := 1
	i := 2
	for got < n {
		i *=2
	}
	return i
}

// method1: iterator all   0(n)
func nthUglyNumberM1(n int) int {
	i := 1
	got := 0
	for got < n {
		if isUgly(i) {
			got++
		}
		if got == n {
			break
		}
		i++
	}
	return i
}
