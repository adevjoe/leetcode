package leetcode

import "testing"

func TestFindPeakElement(t *testing.T) {
	cases := []struct {
		desc string
		nums []int
		want []int
	}{
		{
			desc: "#1",
			nums: []int{1},
			want: []int{0},
		},
		{
			desc: "#2",
			nums: []int{1, 2, 3, 1},
			want: []int{2},
		},
		{
			desc: "#3",
			nums: []int{1, 2, 1, 3, 5, 6, 4},
			want: []int{1, 5},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := findPeakElement(s.nums); !contains(s.want, got) {
				t.Error()
			}
		})
	}
}

func contains(nums []int, int2 int) bool {
	for _, num := range nums {
		if num == int2 {
			return true
		}
	}
	return false
}
