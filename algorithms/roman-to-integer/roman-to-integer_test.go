package leetcode

import "testing"

func TestRomantoInteger(t *testing.T) {
	romanToInt("III")
	romanToInt("IV")
	romanToInt("IX")
	romanToInt("LVIII")
	romanToInt("MCMXCIV")
}
