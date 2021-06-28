// https://leetcode.com/problems/decode-ways

package leetcode

func numDecodings(s string) int {
	cache := make(map[string]int, 0)
	if s[0] == '0' {
		return 0
	}
	cache[""] = 1
	cache[s[:1]] = 1
	for i := 1; i < len(s); i++ {
		if s[i] == '0' {
			if s[i-1] != '1' && s[i-1] != '2' {
				return 0
			}
			cache[s[:i+1]] = cache[s[:i-1]]
			continue
		}
		if s[i-1] == '0' {
			cache[s[:i+1]] = cache[s[:i]]
			continue
		}
		if s[i-1:i+1] > "10" && s[i-1:i+1] <= "26" {
			cache[s[:i+1]] = cache[s[:i]] + cache[s[:i-1]]
			continue
		}
		cache[s[:i+1]] = cache[s[:i]]
	}
	return cache[s]
}

