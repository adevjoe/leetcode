package leetcode

import "testing"

func TestBestTimetoBuyandSellStock(t *testing.T) {
	cases := []struct {
		desc   string
		prices []int
		want   int
	}{
		{
			desc:   "#1",
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
		{
			desc:   "#2",
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
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
