package algorithms

import (
	"testing"
)

func TestIsUglyNumber(t *testing.T) {
	cases := []struct {
		desc string
		num  int
		want bool
	}{
		{desc: "0", num: 0, want: false},
		{desc: "1", num: 1, want: true},
		{desc: "2", num: 2, want: true},
		{desc: "3", num: 3, want: true},
		{desc: "4", num: 4, want: true},
		{desc: "5", num: 5, want: true},
		{desc: "6", num: 6, want: true},
		{desc: "7", num: 7, want: false},
		{desc: "10", num: 10, want: true},
		{desc: "11", num: 11, want: false},
		{desc: "999", num: 999, want: false},
		{desc: "22", num: 22, want: false},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := isUgly(s.num); got != s.want {
				t.Errorf("num: %d, got: %v, want: %v", s.num, got, s.want)
			}
		})
	}
}
