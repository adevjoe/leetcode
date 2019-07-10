/**
Given n non-negative integers a1, a2, ..., an ,
where each represents a point at coordinate (i, ai).
n vertical lines are drawn such that the two endpoints of line i is at (i, ai)
and (i, 0). Find two lines, which together with x-axis forms a container,
such that the container contains the most water.

Like this picture.
https://s3-lc-upload.s3.amazonaws.com/uploads/2018/07/17/question_11.jpg

Note: You may not slant the container and n is at least 2.

Example:

Input: [1,8,6,2,5,4,8,3,7]
Output: 49
*/
package main

import "fmt"

func maxArea(height []int) int {
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

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // 49
	fmt.Println(maxArea([]int{2, 3, 4, 5, 18, 17, 6}))     // 17
	fmt.Println(maxArea([]int{1, 1}))                      // 1
}
