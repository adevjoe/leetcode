package leetcode

import (
	"fmt"
	"testing"
)

func TestCountandSay(t *testing.T) {
	fmt.Println(byte(1))
	cases := []struct {
		desc string
		n    int
		want string
	}{
		{
			desc: "base",
			n:    1,
			want: "1",
		},
		{
			desc: "2",
			n:    2,
			want: "11",
		},
		{
			desc: "3",
			n:    3,
			want: "21",
		},
		{
			desc: "4",
			n:    4,
			want: "1211",
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := countAndSay(s.n); got != s.want {
				t.Errorf("n: %d, want: %s, got: %s", s.n, s.want, got)
			}
		})
	}
}
