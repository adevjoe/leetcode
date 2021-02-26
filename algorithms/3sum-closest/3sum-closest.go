// https://leetcode.com/problems/3sum-closest

package leetcode

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	result := nums[0] + nums[1] + nums[len(nums)-1]
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		start := i + 1
		end := len(nums) - 1
		for start < end {
			sum := nums[i] + nums[start] + nums[end]
			if sum > target {
				end--
			} else {
				start++
			}
			if math.Abs(float64(sum-target)) < math.Abs(float64(result-target)) {
				result = sum
			}
		}
	}
	return result
}
