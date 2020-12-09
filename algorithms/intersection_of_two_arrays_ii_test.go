package algorithms

import (
	"reflect"
	"testing"
)

func TestIntersectionII(t *testing.T) {
	cases := []struct {
		desc  string
		nums1 []int
		nums2 []int
		want  []int
	}{
		{
			desc:  "#1",
			nums1: []int{1, 2, 2, 1},
			nums2: []int{2, 2},
			want:  []int{2, 2},
		},
		{
			desc:  "#2",
			nums1: []int{4, 9, 5},
			nums2: []int{9, 4, 9, 8, 4},
			want:  []int{4, 9},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := intersectionII(s.nums1, s.nums2); !reflect.DeepEqual(s.want, got) {
				t.Errorf("nums1: %v, nums2: %v, want: %v, got: %v", s.nums1, s.nums2, s.want, got)
			}
		})
	}
}
