// https://leetcode.com/problems/trapping-rain-water

package leetcode

func trap(height []int) int {
	left := 0
	right := 0
	isTrap := false
	cur := 0
	rain := 0
	for left < len(height) && right < len(height) {
		if height[left] < 1 {
			left++
			continue
		}
		if right <= left+1 {
			if height[right] > height[left] {
				left = right
			}
			right++
			continue
		}
		if height[right] >= height[left] {
			rain += getArea(height[left : right+1])
			left = right
			isTrap = false
		} else {
			if !isTrap {
				isTrap = true
				cur = right
			}
			if height[right] > height[cur] {
				cur = right
			}
			if isTrap && right == len(height)-1 {
				rain += getArea(height[left : cur+1])
				isTrap = false
				left = cur
				right = cur
			}
			right++
		}

	}
	return rain
}

func getArea(area []int) int {
	min := 0
	rain := 0
	if area[0] > area[len(area)-1] {
		min = len(area) - 1
	}
	for i := 1; i < len(area)-1; i++ {
		if got := area[min] - area[i]; got > 0 {
			rain += got
		}
	}
	return rain
}

// O(n)
func trap2(height []int) int {
	left := 0
	right := len(height) - 1
	maxleft := 0
	maxright := 0
	rain := 0
	for left <= right {
		if height[left] <= height[right] {
			if height[left] >= maxleft {
				maxleft = height[left]
			}
			rain += maxleft - height[left]
			left++
		} else {
			if height[right] >= maxright {
				maxright = height[right]
			}
			rain += maxright - height[right]
			right--
		}
	}
	return rain
}
