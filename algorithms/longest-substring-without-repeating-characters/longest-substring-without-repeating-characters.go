// https://leetcode.com/problems/longest-substring-without-repeating-characters

package leetcode

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	longestSubString := ""
	currentString := ""
	hash := map[int]int{}
	for key, value := range s {
		if key == 0 {
			hash = map[int]int{int(value): 1}
			currentString = string(value)
			longestSubString = currentString
			continue
		}
		if hash[int(value)] > 0 { // repeated
			currentString = s[hash[int(value)] : key+1]
			start := hash[int(value)]
			hash = map[int]int{}
			for i := start; i < key+1; i++ {
				hash[int(s[i])] = i + 1
			}
		} else {
			hash[int(value)] = key + 1
			currentString += string(value)
		}
		if len(currentString) > len(longestSubString) {
			longestSubString = currentString
		}
	}
	return len(longestSubString)
}

func lengthOfLongestSubstring2(s string) int {
	if len(s) == 0 {
		return 0
	}

	max := 1
	cur := 1
	start := 0

	for i := 1; i < len(s); i++ {
		c := s[i]

		// check repeated char
		found := false
		for j := i - 1; j >= start; j-- {
			if s[j] == c {
				cur = i - j
				start = j + 1
				found = true
				break
			}
		}

		// step
		if !found {
			cur++
		}
		if cur > max {
			max = cur
		}
	}

	return max
}
