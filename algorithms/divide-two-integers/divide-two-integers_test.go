package leetcode

import (
	"fmt"
	"testing"
)

func TestDivideTwoIntegers(t *testing.T) {
	fmt.Println(3<<2)
	fmt.Println(1>>1)
	cases := []struct {
		desc     string
		dividend int
		divisor  int
		want     int
	}{
		{
			desc:     "#1",
			dividend: 10,
			divisor:  3,
			want:     3,
		},
		{
			desc:     "#2",
			dividend: 7,
			divisor:  -3,
			want:     -2,
		},
		{
			desc:     "#3",
			dividend: 0,
			divisor:  1,
			want:     0,
		},
		{
			desc:     "#4",
			dividend: 1,
			divisor:  1,
			want:     1,
		},
		{
			desc:     "#5",
			dividend: -1 << 31,
			divisor:  -1,
			want:     1<<31 - 1,
		},
		{
			desc:     "#6",
			dividend: -13,
			divisor:  5,
			want:     -2,
		},
		{
			desc:     "#7",
			dividend: 13,
			divisor:  -5,
			want:     -2,
		},
		{
			desc:     "#8",
			dividend: -13,
			divisor:  -5,
			want:     2,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := divide(s.dividend, s.divisor); s.want != got {
				t.Errorf("dividend: %d, divisor: %d, want: %d, got: %d", s.dividend, s.divisor, s.want, got)
			}
		})
	}
}
