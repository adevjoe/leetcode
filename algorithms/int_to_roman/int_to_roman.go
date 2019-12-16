// https://leetcode.com/problems/integer-to-roman/

package int_to_roman

var RomanList = []struct {
	i     int
	roman string
}{
	{1, "I"},
	{5, "V"},
	{10, "X"},
	{50, "L"},
	{100, "C"},
	{500, "D"},
	{1000, "M"},
}

func intToRoman(num int) string {
	s := ""
	i := len(RomanList) - 1
	for i >= 0 && num > 0 {
		t := 3
		for t >= 0 {
			var tmp, cut int
			var tmpRoman string
			if t > 0 {
				tmp = RomanList[i].i * t
				cut = RomanList[i].i
				tmpRoman += RomanList[i].roman
			}
			if t == 0 {
				if i >= 2 && i%2 == 0 {
					tmp = RomanList[i].i - RomanList[i-2].i
					cut = tmp
					tmpRoman += RomanList[i-2].roman + RomanList[i].roman
				}
				if i >= 1 && i%2 == 1 {
					tmp = RomanList[i].i - RomanList[i-1].i
					cut = tmp
					tmpRoman += RomanList[i-1].roman + RomanList[i].roman
				}
			}
			if num >= tmp {
				num -= cut
				s += tmpRoman
				//fmt.Printf("num: %d, i: %d, t: %d, tmp: %d\n",
				//	num, i, t, tmp)
			}
			t--
		}
		i--
	}
	return s
}

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
