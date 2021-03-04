// https://leetcode.com/problems/house-robber

package leetcode

import "math"

var m map[int]*int

func rob(nums []int) int {
	// 清空 map
	m = map[int]*int{}
	return robSub(nums, 0)
}

func robSub(nums []int, i int) int {
	// 边界检测
	if len(nums[i:]) == 0 {
		return 0
	}
	if len(nums[i:]) == 1 {
		return nums[i]
	}
	if len(nums[i:]) == 2 {
		return int(math.Max(float64(nums[i]), float64(nums[i+1])))
	}

	// 每次分成两种解法，求两种解法的最大值
	/**
	1. nums[i] +  nums[i+2]
	2. nums[i+1] +  nums[i+3]
	max(1, 2)
	*/

	// l 解法1，r 解法2
	var l, r int
	// 解法 1 计算
	if m[i+2] != nil && *m[i+2] > -1 { // 存储已有的解法，避免计算相同的解法
		l = *m[i+2]
	} else {
		l = robSub(nums, i+2)
		m[i+2] = &l
	}

	// 解法 2 计算
	if len(nums[i:]) >= 4 {
		if m[i+3] != nil && *m[i+3] > -1 {
			r = *m[i+3]
		} else {
			r = robSub(nums, i+3)
			m[i+3] = &r
		}
	}

	// 计算两个解法的最大值
	return int(math.Max(float64(l+nums[i]), float64(r+nums[i+1])))
}
