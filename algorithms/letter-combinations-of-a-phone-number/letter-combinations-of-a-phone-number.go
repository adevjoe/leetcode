// https://leetcode.com/problems/letter-combinations-of-a-phone-number

package leetcode

var letterMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	var result []string
	for _, s := range digits {
		result = combine(result, letterMap[string(s)])
	}
	return result
}

func combine(a []string, b string) []string {
	var result []string
	if len(a) == 0 {
		for _, v := range b {
			result = append(result, string(v))
		}
	} else {
		for _, av := range a {
			for _, bv := range b {
				result = append(result, av+string(bv))
			}
		}
	}
	return result
}
