// https://leetcode.com/problems/remove-element

package leetcode

func removeElement(nums []int, val int) int {
	var n int
	for i := range nums {
		if nums[i] != val {
			nums[n] = nums[i]
			n++
		}
	}
	return n
}
