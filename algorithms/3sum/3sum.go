// https://leetcode.com/problems/3sum

package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			lo := i + 1
			hi := len(nums) - 1
			sum := 0 - nums[i]
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
