// https://leetcode.com/problems/jump-game

package leetcode

func canJump(nums []int) bool {
	for i := range nums {
		if nums[i] == 0 {
			jump := false
			tmplast := i
			for j := i - 1; j >= 0; j-- {
				plus := 0
				if nums[tmplast] == 0 {
					plus = 1
				}
				if nums[j]+j < tmplast+plus {
					jump = false
				} else {
					jump = true
					tmplast = j
				}
			}
			if !jump && i != len(nums)-1 {
				return false
			}
		}
	}
	return true
}

func canJump2(nums []int) bool {
	last := len(nums) - 1
	for i := last - 1; i >= 0; i-- {
		if i+nums[i] >= last {
			last = i
		}
	}
	return last <= 0
}
