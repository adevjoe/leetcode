// https://leetcode.com/problems/majority-element

package leetcode

func majorityElement(nums []int) int {
	cur := 1
	last := len(nums) - 1
	for cur < last {
		if nums[cur] != nums[cur-1] {
			nums[cur], nums[last] = nums[last], nums[cur]
			last--
		} else {
			cur++
		}
	}
	return nums[(len(nums)-1)/2]
}
