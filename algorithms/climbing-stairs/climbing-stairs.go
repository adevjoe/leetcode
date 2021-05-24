// https://leetcode.com/problems/climbing-stairs

package leetcode

func climbStairs(n int) int {
	s := make([]int, 3)
	s[1] = 1
	s[2] = 2
	for i := 3; i <= n; i++ {
		s = append(s, s[i-1]+s[i-2])
	}
	return s[n]
}

func climbStairs2(n int) int {
	if n < 3 {
		return n
	}
	lastlast := 1
	last := 2
	cur := 0
	for i := 3; i <= n; i++ {
		cur = lastlast + last
		lastlast = last
		last = cur
	}
	return cur
}
