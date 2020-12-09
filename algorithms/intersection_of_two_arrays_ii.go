// https://leetcode.com/problems/intersection-of-two-arrays-ii/

package algorithms

import "sort"

func intersectionII(nums1 []int, nums2 []int) []int {
	if nums1 == nil || nums2 == nil {
		return nil
	}
	sort.Ints(nums1)
	sort.Ints(nums2)
	var a, b int
	var result []int
	for a < len(nums1) && b < len(nums2) {
		if nums1[a] == nums2[b] {
			result = append(result, nums1[a])
			a++
			b++
			continue
		}
		if nums1[a] > nums2[b] {
			b++
		} else {
			a++
		}
	}
	return result
}
