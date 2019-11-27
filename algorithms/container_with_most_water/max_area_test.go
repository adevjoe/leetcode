package container_with_most_water

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {
	fmt.Println(MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // 49
	fmt.Println(MaxArea([]int{2, 3, 4, 5, 18, 17, 6}))     // 17
	fmt.Println(MaxArea([]int{1, 1}))                      // 1
}
