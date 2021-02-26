package leetcode

import (
	"fmt"
	"testing"
)

func TestContainerWithMostWater(t *testing.T) {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // 49
	fmt.Println(maxArea([]int{2, 3, 4, 5, 18, 17, 6}))     // 17
	fmt.Println(maxArea([]int{1, 1}))                      // 1
}
