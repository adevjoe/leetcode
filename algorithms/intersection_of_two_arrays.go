// https://leetcode.com/problems/intersection-of-two-arrays/

package algorithms

import "sort"

func intersection(nums1 []int, nums2 []int) []int {
	if nums1 == nil || nums2 == nil {
		return nil
	}
	sort.Sort(sort.Reverse(sort.IntSlice(nums1)))
	sort.Sort(sort.Reverse(sort.IntSlice(nums2)))
	var a, b int
	var result []int
	for a < len(nums1) && b < len(nums2) {
		if len(result) > 0 && nums1[a] == result[len(result)-1] {
			a++
			continue
		}
		if len(result) > 0 && nums2[b] == result[len(result)-1] {
			b++
			continue
		}
		if nums1[a] == nums2[b] {
			result = append(result, nums1[a])
			a++
			b++
			continue
		}
		if nums1[a] > nums2[b] {
			a++
		} else {
			b++
		}
	}
	return result
}
