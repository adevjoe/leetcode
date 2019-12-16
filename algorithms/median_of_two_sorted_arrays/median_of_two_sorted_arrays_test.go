package median_of_two_sorted_arrays

import (
	"fmt"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	fmt.Println(findMedianSortedArrays([]int{1, 2, 3, 4}, []int{4, 10000}))
}
