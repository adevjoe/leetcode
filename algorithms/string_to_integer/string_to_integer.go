// https://leetcode.com/problems/string-to-integer-atoi/

package string_to_integer

import (
	"math"
)

// 整理为四个规则 1.检测空白 2.排除字母 3.记录符号 4.避免溢出
// 优化改进，执行效率提升1倍
// 之前连字母也会检验，现在不是空白符、数字和符号的直接 break，起到提前退出程序的效果，减少程序运行次数
func Atoi(str string) int {
	num := 0
	sign := ""
	start := false // 一旦解析到数字或符号则为 true
	for key := range str {
		// 校验空白符
		if str[key:key+1] == " " && !start {
			continue
		}
		// 校验符号
		if (str[key:key+1] == "+" || str[key:key+1] == "-") && !start {
			start = true
			sign = str[key : key+1]
			continue
		}
		// 校验数字
		if str[key:key+1] >= "0" && str[key:key+1] <= "9" {
			// 检测溢出
			if num > math.MaxInt32/10 || (num == math.MaxInt32/10 && str[key:key+1] > "7") {
				if sign == "-" {
					return math.MinInt32
				}
				return math.MaxInt32
			}
			start = true
			num = num*10 + int(str[key]) - 48
			continue
		}
		// 不符合上述规则的直接结束
		break
	}
	if sign == "-" {
		num = -num
	}
	return num
}

// 初次提交
func myAtoi(str string) int {
	i := 0
	sign := 0 // 0 无符号 / 1 + / 2 -
	hasNum := false
	a := []byte(str)
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
