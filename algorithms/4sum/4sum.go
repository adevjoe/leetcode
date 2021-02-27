// https://leetcode.com/problems/4sum

package leetcode

import "sort"

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return [][]int{}
	}
	sort.Ints(nums)
	var result [][]int
	for i := 0; i <= len(nums)-4; i++ {
		if i > 0 && nums[i] == nums[i-1] { // avoid duplicate
			continue
		}
		a := threeSum(nums[i+1:], target-nums[i])
		for _, r := range a {
			b := []int{nums[i]}
			b = append(b, r...)
			result = append(result, b)
		}
	}
	return result
}

func threeSum(nums []int, target int) [][]int {
	var result [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			lo := i + 1
			hi := len(nums) - 1
			sum := target - nums[i]
			for lo < hi {
				if nums[lo]+nums[hi] == sum {
					result = append(result, []int{nums[i], nums[lo], nums[hi]})
					for lo < hi && nums[lo] == nums[lo+1] {
						lo++
					}
					for lo < hi && nums[hi] == nums[hi-1] {
						hi--
					}
					lo++
					hi--
				} else if nums[lo]+nums[hi] < sum {
					lo++
				} else {
					hi--
				}
			}
		}
	}
	return result
}
