// https://leetcode.com/problems/sqrtx

package leetcode

func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	left := 0
	right := x
	for left < right {
		mid := (left + right) / 2
		mm := mid * mid
		if mm == x {
			return mid
		}
		if mm > x {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left - 1
}
