package leetcode

import "testing"

func TestSingleNumber(t *testing.T) {
	cases := []struct {
		desc string
		nums []int
		want int
	}{
		{
			desc: "#1",
			nums: []int{1},
			want: 1,
		},
		{
			desc: "#2",
			nums: []int{2, 2, 1},
			want: 1,
		},
		{
			desc: "#3",
			nums: []int{4, 1, 2, 1, 2},
			want: 4,
		},
		{
			desc: "#4",
			nums: []int{1, 4, 1, 2, 2},
			want: 4,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := singleNumber(s.nums); got != s.want {
				t.Error()
			}
		})
	}
}
