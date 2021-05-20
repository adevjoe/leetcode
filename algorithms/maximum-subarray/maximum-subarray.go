// https://leetcode.com/problems/maximum-subarray

package leetcode

func maxSubArray(nums []int) int {
	bestSum := nums[0]
	curSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if curSum <= 0 {
			curSum = nums[i]
		} else {
			curSum += nums[i]
		}
		if curSum > bestSum {
			bestSum = curSum
		}
	}
	return bestSum
}
