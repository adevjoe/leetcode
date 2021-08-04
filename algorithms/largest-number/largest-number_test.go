package leetcode

import "testing"

func TestLargestNumber(t *testing.T) {
	cases := []struct {
		desc string
		nums []int
		want string
	}{
		{
			desc: "#1",
			nums: []int{0, 0},
			want: "0",
		},
		{
			desc: "#2",
			nums: []int{10, 2},
			want: "210",
		},
		{
			desc: "#3",
			nums: []int{3, 30, 34, 5, 9},
			want: "9534330",
		},
		{
			desc: "#4",
			nums: []int{1},
			want: "1",
		},
		{
			desc: "#5",
			nums: []int{10},
			want: "10",
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := largestNumber(s.nums); got != s.want {
				t.Error()
			}
		})
	}
}
