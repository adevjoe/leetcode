// https://leetcode.com/problems/permutations

package leetcode

func permute(nums []int) [][]int {
	var result [][]int
	if len(nums) == 1 {
		return [][]int{nums}
	}
	for i := range nums {
		var tmpnums []int
		tmpnums = append(tmpnums, nums[:i]...)
		tmpnums = append(tmpnums, nums[i+1:]...)
		for _, n := range permute(tmpnums) {
			result = append(result, append([]int{nums[i]}, n...))
		}
	}
	return result
}
