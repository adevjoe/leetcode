/**
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

Example 1:

nums1 = [1, 3]
nums2 = [2]

The median is 2.0

Example 2:

nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5
 */
package main

import "fmt"

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2, 3, 4}, []int{4,10000}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	num := make([]int, len1+len2, len1+len2)
	t := (len1 + len2) / 2

	var i, j, k int
	for i < len1 && j < len2 && k <= t {
		if nums1[i] <= nums2[j] {
			num[k] = nums1[i]
			i ++
		} else {
			num[k] = nums2[j]
			j ++
		}
		k ++
		fmt.Println(i, j, num)
	}

	for i < len1 && k <= t {
		num[k] = nums1[i]
		i ++
		k ++
		fmt.Println(i, j, num)
	}

	for j < len2 && k <= t {
		num[k] = nums2[j]
		j ++
		k ++
		fmt.Println(i, j, num)
	}

	if t*2 < (len1 + len2) {
		return float64(num[t])
	}
	return float64(num[t-1]+num[t]) / 2
}
