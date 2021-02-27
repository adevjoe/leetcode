// https://leetcode.com/problems/unique-binary-search-trees

package leetcode

/**
n = 1
1

n = 2
  1         2
    2     1

n = 3

  1         1          2            3         3
     2         3     1   3       2         1
        3    2                 1             2
*/
func numTrees(n int) int {
	if n == 0 {
		return 1
	}
	return numTrees2(1, n)
}

func numTrees2(start, end int) int {
	var nums int
	if start > end {
		return 0
	}

	if start == end {
		return 1
	}

	for i := start; i <= end; i++ {
		var temp int
		left := numTrees2(start, i-1)
		right := numTrees2(i+1, end)
		if left == 0 || right == 0 {
			temp = left + right
		} else {
			temp = left * right
		}
		nums += temp
	}
	return nums
}

// G(n) = G(0) * G(n-1) + G(1) * G(n-2) + … + G(n-1) * G(0)
func numTrees3(n int) int {
	if n == 0 {
		return 1
	}
	nums := make([]int, n+1)
	nums[0], nums[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			nums[i] += nums[j-1] * nums[i-j]
		}
	}
	return nums[n]
}
