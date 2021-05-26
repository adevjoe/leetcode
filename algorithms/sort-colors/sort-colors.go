// https://leetcode.com/problems/sort-colors

package leetcode

func sortColors(nums []int) {
	if len(nums) == 1 {
		return
	}
	left := 0
	right := len(nums) - 1
	i := 0
	for i <= right {
		switch nums[i] {
		case 0:
			if i == left {
				left++
				i++
				break
			}
			nums[i], nums[left] = nums[left], nums[i]
			left++
			break
		case 1:
			i++
			break
		case 2:
			nums[i], nums[right] = nums[right], nums[i]
			right--
			break
		}
	}
}
