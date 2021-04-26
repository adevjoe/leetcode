package leetcode

import "testing"

func TestTrappingRainWater(t *testing.T) {
	cases := []struct {
		desc   string
		height []int
		want   int
	}{
		{
			desc:   "#1",
			height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:   6,
		},
		{
			desc:   "#2",
			height: []int{4, 2, 0, 3, 2, 5},
			want:   9,
		},
		{
			desc:   "#3",
			height: []int{0, 1, 2, 0, 3, 0, 1, 2, 0, 0, 4, 2, 1, 2, 5, 0, 1, 2, 0, 2},
			want:   26,
		},
		{
			desc:   "#4",
			height: []int{5, 4, 1, 2},
			want:   1,
		},
		{
			desc:   "#5",
			height: []int{2, 6, 3, 8, 2, 7, 2, 5, 0},
			want:   11,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := trap(s.height); got != s.want {
				t.Errorf("height: %v, want: %d, got: %d", s.height, s.want, got)
			}
			if got := trap2(s.height); got != s.want {
				t.Errorf("height: %v, want: %d, got: %d", s.height, s.want, got)
			}
		})
	}
}
