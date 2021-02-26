// https://leetcode.com/problems/roman-to-integer

package leetcode

var RomanMap = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50,
	"C": 100, "D": 500, "M": 1000}

func romanToInt(s string) int {
	i := 0
	num := 0
	for i <= len(s)-1 {
		if i == len(s)-1 {
			num += RomanMap[string(s[i])]
		} else if RomanMap[string(s[i])] >= RomanMap[string(s[i+1])] {
			num += RomanMap[string(s[i])]
		} else {
			num += RomanMap[string(s[i+1])] - RomanMap[string(s[i])]
			i++
		}
		i++
	}
	return num
}
