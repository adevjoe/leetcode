// https://leetcode.com/problems/largest-number

package leetcode

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	var s []string
	for i := range nums {
		s = append(s, strconv.Itoa(nums[i]))
	}

	sort.Sort(numsSlice(s))
	if s[0][0] == '0' {
		return "0"
	}
	result := ""
	for i := range s {
		result += s[i]
	}
	return result
}

type numsSlice []string

func (n numsSlice) Len() int {
	return len(n)
}

func (n numsSlice) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n numsSlice) Less(i, j int) bool {
	return n[i]+n[j] > n[j]+n[i]
}
