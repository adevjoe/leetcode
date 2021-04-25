// https://leetcode.com/problems/first-missing-positive

package leetcode

import "sort"

// place num one their own area
func firstMissingPositive(nums []int) int {
	for i := range nums {
		for nums[i] > 0 && nums[i] <= len(nums) && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i, num := range nums {
		if num != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func firstMissingPositive2(nums []int) int {
	min := 1
	sort.Ints(nums)
	for _, num := range nums {
		if num <= 0 {
			continue
		}
		if num > min {
			break
		}
		if num == min {
			min++
		}
	}
	return min
}
