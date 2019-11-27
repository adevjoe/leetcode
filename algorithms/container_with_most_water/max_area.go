// https://leetcode.com/problems/container-with-most-water/

package container_with_most_water

func MaxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := 0
	for left != right {
		m := minInt(height[left], height[right]) * (right - left)
		if m > maxArea {
			maxArea = m
		}
		if height[left] > height[right] {
			//a := height[right]
			right--
			//for (right) > left && height[right] <= a {
			//	right--
			//}
		} else {
			//b := height[left]
			left++
			//for (left) < right && height[left] <= b {
			//	left++
			//}
		}
	}
	return maxArea
}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
