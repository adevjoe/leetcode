// https://leetcode.com/problems/palindrome-partitioning

package leetcode

var cache map[string]bool

func partition(s string) [][]string {
	cache = make(map[string]bool)
	return partitionPlus(s)
}

func partitionPlus(s string) [][]string {
	result := make([][]string, 0)
	if len(s) == 1 {
		result = append(result, []string{s})
		return result
	}
	for j := 0; j < len(s); j++ {
		if s[0] == s[j] && (j <= 2 || cache[s[1:j]]) {
			cache[s[:j+1]] = true
			tail := partitionPlus(s[j+1:])
			for k := range tail {
				result = append(result, append([]string{s[:j+1]}, tail[k]...))
			}
			if j == len(s)-1 {
				result = append(result, []string{s})
			}
		}
	}
	return result
}
