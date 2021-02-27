package leetcode

import "testing"

func TestRemoveElement(t *testing.T) {
	cases := []struct {
		desc string
		nums []int
		val  int
		want int
	}{
		{
			desc: "#1",
			nums: []int{3, 2, 2, 3},
			val:  3,
			want: 2,
		},
		{
			desc: "#2",
			nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:  2,
			want: 5,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := removeElement(s.nums, s.val); got != s.want {
				t.Errorf("nums: %v, val: %d, want: %d, got: %d", s.nums, s.val, s.want, got)
			}
		})
	}
}
