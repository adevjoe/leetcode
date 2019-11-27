// https://leetcode.com/problems/two-sum/

package two_sum

func TwoSum(nums []int, target int) []int {
	for key, value := range nums {
		if (len(nums) - key) > 0 {
			for i := key + 1; i < len(nums); i++ {
				if (value + nums[i]) == target {
					return []int{key, i}
				}
			}
		}
	}
	return []int{}
}
