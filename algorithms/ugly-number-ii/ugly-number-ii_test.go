package leetcode

import "testing"

func TestUglyNumberII(t *testing.T) {
	cases := []struct {
		desc string
		num  int
		want int
	}{
		{desc: "1", num: 1, want: 1},
		{desc: "2", num: 2, want: 2},
		{desc: "3", num: 3, want: 3},
		{desc: "4", num: 4, want: 4},
		{desc: "5", num: 5, want: 5},
		{desc: "7", num: 7, want: 8},
		{desc: "10", num: 10, want: 12},
		{desc: "1352", num: 1352, want: 402653184},
		{desc: "2123366400", num: 1690, want: 2123366400},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := nthUglyNumber(s.num); got != s.want {
				t.Errorf("num: %d, got: %v, want: %v", s.num, got, s.want)
			}
		})
	}
}
