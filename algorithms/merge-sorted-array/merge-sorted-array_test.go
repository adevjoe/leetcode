package leetcode

import (
	"reflect"
	"testing"
)

func TestMergeSortedArray(t *testing.T) {
	cases := []struct {
		desc  string
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{
			desc:  "#1",
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
			want:  []int{1, 2, 2, 3, 5, 6},
		},
		{
			desc:  "#2",
			nums1: []int{7, 8, 0, 0, 0},
			m:     2,
			nums2: []int{2, 5, 6},
			n:     3,
			want:  []int{2, 5, 6, 7, 8},
		},
		{
			desc:  "#3",
			nums1: []int{1},
			m:     1,
			nums2: []int{},
			n:     0,
			want:  []int{1},
		},
		{
			desc:  "#4",
			nums1: []int{0},
			m:     0,
			nums2: []int{1},
			n:     1,
			want:  []int{1},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if merge(s.nums1, s.m, s.nums2, s.n); !reflect.DeepEqual(s.nums1, s.want) {
				t.Error()
			}
		})
	}
}
