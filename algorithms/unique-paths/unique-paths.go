// https://leetcode.com/problems/unique-paths

package leetcode

func uniquePaths(m int, n int) int {
	for m > 1 && n > 1 {
		down := uniquePaths(m-1, n)
		right := uniquePaths(m, n-1)
		return down + right
	}
	return 1
}

func uniquePaths2(m int, n int) int {
	if m < n {
		m, n = n, m
	}
	cur := make(map[int]int)
	for i := 0; i <= m; i++ {
		cur[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			cur[j] += cur[j-1]
		}
	}
	return cur[n-1]
}
