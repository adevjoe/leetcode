// https://leetcode.com/problems/maximum-product-subarray

package leetcode

func maxProduct(nums []int) int {
	max := nums[0]
	imax := max
	imin := max
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			imax, imin = imin, imax
		}
		imax = Max(nums[i], imax*nums[i])
		imin = Min(nums[i], imin*nums[i])
		max = Max(max, imax)
	}
	return max
}

func Max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func Min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}
