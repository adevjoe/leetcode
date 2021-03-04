// https://leetcode.com/problems/ugly-number-ii

package leetcode

import "math"

// 每个丑数由任意个 2，3，5 相乘组成，因此可以变成 ?*2, ?*3, ?*5 来构成，? 为之前的丑数。
// 我们可以记住每个丑数的位置，计算 ?*2, ?*3, ?*5 的最小值，由此得到当前的丑数。
func nthUglyNumber(n int) int {
	t2, t3, t5 := 1, 1, 1
	r := map[int]int{}
	r[1] = 1
	for i := 2; i <= n; i++ {
		r[i] = int(math.Min(math.Min(float64(r[t2]*2), float64(r[t3]*3)), float64(r[t5]*5)))
		if r[i] == r[t2]*2 {
			t2++
		}
		if r[i] == r[t3]*3 {
			t3++
		}
		if r[i] == r[t5]*5 {
			t5++
		}
	}
	return r[n]
}
