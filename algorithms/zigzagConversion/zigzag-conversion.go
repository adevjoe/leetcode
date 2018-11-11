/**
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this:
(you may want to display this pattern in a fixed font for better legibility)

https://leetcode.com/problems/zigzag-conversion/description/

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:

P     I    N
A   L S  I G
Y A   H R
P     I
*/
package zigzagConversion

import "fmt"

func main() {
	fmt.Println(splitGroup("ABCDE", 4))
	fmt.Println(convert("ABCDE", 4))
}

func convert(s string, numRows int) string {
	if s == "" {
		return ""
	}
	groupsNum, _, groups, _ := splitGroup(s, numRows)
	if groupsNum == 1 && (len(s) <= numRows) {
		return s
	}
	rows := make([]string, numRows)
	for _, str := range groups {
		for key, value := range str {
			if key+1 <= numRows {
				rows[key] += string(value)
			} else {
				rows[2*numRows-key-2] += string(value)
			}
		}
	}
	ss := ""
	for _, value := range rows {
		ss += value
	}
	return ss
}

// 字符串分组
func splitGroup(s string, numRows int) (groupsNum int, groupPreNum int, groups []string, remain string) {
	l := len(s)
	if numRows == 1 {
		groupPreNum = 1
	} else {
		groupPreNum = numRows*2 - 2
	}
	groupsNum = l / groupPreNum
	groups = make([]string, groupsNum)
	for i := 0; i <= groupsNum-1; i++ {
		groups[i] = s[i*groupPreNum : (i+1)*groupPreNum]
	}
	remain = s[groupPreNum*groupsNum : l]
	groups = append(groups, remain)
	if len(remain) > 0 {
		groupsNum++
	}
	return
}
