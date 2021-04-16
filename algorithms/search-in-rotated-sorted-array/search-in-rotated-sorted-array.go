// https://leetcode.com/problems/search-in-rotated-sorted-array

package leetcode

// O(log n)
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for right-left > 1 {
		mid := getMidIndex(left, right)
		if nums[mid] == target {
			return mid
		}
		if nums[left] == target {
			return left
		}
		if nums[right] == target {
			return right
		}
		if nums[left] < nums[mid] {
			if target < nums[mid] && target > nums[left] {
				right = mid
				continue
			} else {
				left = mid
				continue
			}
		}
		if nums[mid] < nums[right] {
			if target > nums[mid] && target < nums[right] {
				left = mid
			} else {
				right = mid
				continue
			}
		}
	}
	if nums[left] == target {
		return left
	}
	if nums[right] == target {
		return right
	}
	return -1
}

func getMidIndex(left, right int) int {
	return (right - left + 1) / 2 + left
}

func searchSub(nums []int, target, from int) int {
	if len(nums) <= 2 {
		for i, v := range nums {
			if v == target {
				return i + from
			}
		}
		return -1
	}
	if nums[len(nums)-1] > nums[0] && (target < nums[0] && target > nums[len(nums)-1]) {
		return -1
	}
	mid := len(nums) / 2
	left := searchSub(nums[:mid+1], target, from)
	if left != -1 {
		return left
	}
	right := searchSub(nums[mid+1:], target, from+mid+1)
	if right != -1 {
		return right
	}

	return -1
}

// O(n)
func search1(nums []int, target int) int {
	for i, v := range nums {
		if v == target {
			return i
		}
	}
	return -1
}
