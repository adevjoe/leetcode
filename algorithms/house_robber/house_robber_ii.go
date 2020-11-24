package house_robber

import "math"

func robTwo(nums []int) int {
	var max int
	for i := range nums {
		var tmp []int
		if len(nums)-3 >= i {
			if i == 0 {
				tmp = append(tmp, nums[i+2:len(nums)-1]...)
			} else {
				tmp = append(tmp, nums[i+2:]...)
			}
		}
		if i >= 2 {
			if i == len(nums)-1 {
				tmp = append(tmp, nums[1:i-1]...)
			} else {
				tmp = append(tmp, nums[:i-1]...)
			}
		}
		max = int(math.Max(float64(max), float64(nums[i]+rob(tmp))))
	}
	return max
}
