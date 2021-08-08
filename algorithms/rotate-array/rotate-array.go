// https://leetcode.com/problems/rotate-array

package leetcode

func rotate(nums []int, k int) {
	l := len(nums)
	if l <= 1 {
		return
	}
	step := k % l
	reverse(nums, 0, l-1)
	reverse(nums, 0, step-1)
	reverse(nums, step, l-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
