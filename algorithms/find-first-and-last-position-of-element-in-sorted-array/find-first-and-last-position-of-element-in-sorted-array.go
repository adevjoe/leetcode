// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array

package leetcode

// log(n)
func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	start, end := -1, -1
	for left < right { // find left
		if right-1 == left {
			if nums[left] == target {
				start = left
				break
			}
			if nums[right] == target {
				start = right
			}
			break
		}
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid
		}
	}
	right = len(nums) - 1
	for left < right { // find right
		if right-1 == left {
			if nums[left] == target {
				end = left
			}
			if nums[right] == target {
				end = right
			}
			break
		}
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid
		} else {
			left = mid
		}
	}
	if len(nums) == 1 && nums[0] == target {
		start, end = 0, 0
	}
	return []int{start, end}
}

// O(n)
func searchRange1(nums []int, target int) []int {
	start, end := -1, -1
	for i, num := range nums {
		if num == target {
			end = i
			if start == -1 {
				start = i
			}
		}
	}
	return []int{start, end}
}
