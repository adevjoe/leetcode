package leetcode

import "testing"

func TestRemoveDuplicatesfromSortedArray(t *testing.T) {
	cases := []struct {
		desc string
		nums []int
		want int
	}{
		{
			desc: "#1",
			nums: []int{1, 1, 2},
			want: 2,
		},
		{
			desc: "#2",
			nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want: 5,
		},
	}
	for _, s := range cases {
		if got := removeDuplicates(s.nums); s.want != got {
			t.Errorf("nums: %v, want: %d, got: %d", s.nums, s.want, got)
		}
	}

	for _, s := range cases {
		if got := removeDuplicates2(s.nums); s.want != got {
			t.Errorf("nums: %v, want: %d, got: %d", s.nums, s.want, got)
		}
	}
}
