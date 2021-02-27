// https://leetcode.com/problems/remove-duplicates-from-sorted-array

package leetcode

func removeDuplicates(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return l
	}
	i := 1
	for i < len(nums) {
		if nums[i] == nums[i-1] {
			// delete operation too many times
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return len(nums)
}

func removeDuplicates2(nums []int) int {
	l := len(nums)
	count := 0
	for i := 1; i < l; i++ {
		if nums[i] == nums[i-1] {
			count++
		} else {
			// modify value instead of delete value
			nums[i-count] = nums[i]
		}
	}
	return l - count
}
