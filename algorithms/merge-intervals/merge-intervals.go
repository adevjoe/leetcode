// https://leetcode.com/problems/merge-intervals

package leetcode

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Sort(intervalSort(intervals))
	left := intervals[0][0]
	right := intervals[0][1]
	result := make([][]int, 0)
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] == right {
			right = intervals[i][1]
			continue
		}
		if intervals[i][0] > right {
			result = append(result, []int{left, right})
			left = intervals[i][0]
			right = intervals[i][1]
			continue
		}
		if intervals[i][1] <= right {
			continue
		}
		if intervals[i][1] > right {
			right = intervals[i][1]
			continue
		}
	}
	result = append(result, []int{left, right})
	return result
}

type intervalSort [][]int

func (n intervalSort) Len() int {
	return len(n)
}

func (n intervalSort) Less(i, j int) bool {
	if n[i][0] < n[j][0] {
		return true
	}
	return false
}

func (n intervalSort) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
