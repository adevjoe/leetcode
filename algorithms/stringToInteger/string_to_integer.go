/**
Implement atoi which converts a string to an integer.

The function first discards as many whitespace characters as necessary until the first non-whitespace character is found. Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as possible, and interprets them as a numerical value.

The string can contain additional characters after those that form the integral number, which are ignored and have no effect on the behavior of this function.

If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists because either str is empty or it contains only whitespace characters, no conversion is performed.

If no valid conversion could be performed, a zero value is returned.

Note:

Only the space character ' ' is considered as whitespace character.
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. If the numerical value is out of the range of representable values, INT_MAX (231 − 1) or INT_MIN (−231) is returned.

Example 1:

Input: "42"
Output: 42

Example 2:

Input: "   -42"
Output: -42
Explanation: The first non-whitespace character is '-', which is the minus sign.
             Then take as many numerical digits as possible, which gets 42.
Example 3:

Input: "4193 with words"
Output: 4193
Explanation: Conversion stops at digit '3' as the next character is not a numerical digit.
Example 4:

Input: "words and 987"
Output: 0
Explanation: The first non-whitespace character is 'w', which is not a numerical
             digit or a +/- sign. Therefore no valid conversion could be performed.
Example 5:

Input: "-91283472332"
Output: -2147483648
Explanation: The number "-91283472332" is out of the range of a 32-bit signed integer.
             Thefore INT_MIN (−231) is returned.
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myAtoi("0-1"))
}

func myAtoi(str string) int {
	i := 0
	sign := 0 // 0 无符号 / 1 + / 2 -
	hasNum := false
	a := []byte(str)
	fmt.Println(a)
	for _, value := range a {
		// 排除字母
		if value < 48 || value > 57 {
			if sign > 0 || hasNum {
				break
			}
			if value != 45 && value != 43 && value != 32 {
				break
			}
		}
		// 记录符号
		if i == 0 {
			if value == 43 {
				sign = 1
				continue
			}
			if value == 45 {
				sign = 2
				continue
			}
		}
		if b := isNumber(value); b != -1 {
			hasNum = true
			i = i*10 + b
			if i > math.MaxInt32 {
				break
			}
		}
	}
	if i > math.MaxInt32 {
		if sign == 2 {
			return math.MinInt32
		} else {
			return math.MaxInt32
		}
	}
	if sign == 2 {
		i = i - 2*i
	}
	return i
}

func isNumber(b byte) int {
	if b >= 48 && b <= 57 {
		return int(b) - 48
	}
	return -1
}
