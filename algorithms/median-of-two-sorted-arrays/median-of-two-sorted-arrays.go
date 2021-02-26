// https://leetcode.com/problems/median-of-two-sorted-arrays

package leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	num := make([]int, len1+len2, len1+len2)
	t := (len1 + len2) / 2

	var i, j, k int
	for i < len1 && j < len2 && k <= t {
		if nums1[i] <= nums2[j] {
			num[k] = nums1[i]
			i++
		} else {
			num[k] = nums2[j]
			j++
		}
		k++
	}

	for i < len1 && k <= t {
		num[k] = nums1[i]
		i++
		k++
	}

	for j < len2 && k <= t {
		num[k] = nums2[j]
		j++
		k++
	}

	if t*2 < (len1 + len2) {
		return float64(num[t])
	}
	return float64(num[t-1]+num[t]) / 2
}
