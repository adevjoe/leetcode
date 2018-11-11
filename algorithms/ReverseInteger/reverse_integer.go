/**
Given a 32-bit signed integer, reverse digits of an integer.

https://leetcode.com/problems/reverse-integer/description/
*/
package ReverseInteger

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(2147483647))
}

func reverse(x int) int {
	f := false
	var y int
	if x < 0 {
		f = true
		x = -x
	}
	var array []int
	for x >= 10 {
		a := x % 10
		x = (x - a) / 10
		array = append(array, a)
	}
	if x > 0 {
		array = append(array, x)
	}
	l := len(array)
	for i := 0; i < l; i++ {
		y = y*10 + array[i]
	}
	if f {
		y = -y
	}
	// 这里使用 math 包的最大最小 int32 的值，能提升一倍速度
	if y > math.MaxInt32 || y < math.MinInt32 {
		return 0
	}
	return y
}
