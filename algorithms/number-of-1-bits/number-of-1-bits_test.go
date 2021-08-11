package leetcode

import "testing"

func TestNumberof1Bits(t *testing.T) {
	cases := []struct {
		desc string
		num  uint32
		want int
	}{
		{
			desc: "#1",
			num:  11,
			want: 3,
		},
		{
			desc: "#2",
			num:  128,
			want: 1,
		},
		{
			desc: "#3",
			num:  4294967293,
			want: 31,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := hammingWeight(s.num); got != s.want {
				t.Error()
			}
		})
	}
}
