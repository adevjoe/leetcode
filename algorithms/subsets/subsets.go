// https://leetcode.com/problems/subsets

package leetcode

func subsets(nums []int) [][]int {
	return append([][]int{{}}, subsets2(nums)...)
}

func subsets2(nums []int) [][]int {
	result := make([][]int, 0)
	result = append(result, []int{nums[0]})
	if len(nums[1:]) == 0 {
		return result
	}
	a := subsets2(nums[1:])
	result = append(result, a...)
	for i := range a {
		result = append(result, append([]int{nums[0]}, a[i]...))
	}
	return result
}
