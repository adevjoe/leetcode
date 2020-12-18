// https://leetcode.com/problems/ugly-number/

package algorithms

func isUgly(num int) bool {
	if num <= 0 {
		return false
	}
	if num == 1 {
		return true
	}
	for _, i := range []int{2, 3, 5} {
		for num%i == 0 {
			num /= i
			if num == 1 {
				return true
			}
		}
	}
	return false
}
