package leetcode

import (
	"testing"
)

func TestReverseBits(t *testing.T) {
	cases := []struct {
		desc string
		num  uint32
		want uint32
	}{
		{
			desc: "#1",
			num:  43261596,
			want: 964176192,
		},
		{
			desc: "#2",
			num:  4294967293,
			want: 3221225471,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := reverseBits(s.num); got != s.want {
				t.Error()
			}
		})
	}
}
