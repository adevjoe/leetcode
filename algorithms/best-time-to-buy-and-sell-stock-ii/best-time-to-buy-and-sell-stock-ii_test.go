package leetcode

import "testing"

func TestBestTimetoBuyandSellStockII(t *testing.T) {
	cases := []struct {
		desc   string
		prices []int
		want   int
	}{
		{
			desc:   "#1",
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   7,
		},
		{
			desc:   "#2",
			prices: []int{1, 2, 3, 4, 5},
			want:   4,
		},
		{
			desc:   "#3",
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
		{
			desc:   "#4",
			prices: []int{2, 1, 2, 0, 1},
			want:   2,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := maxProfit(s.prices); got != s.want {
				t.Error()
			}
		})
	}
}
