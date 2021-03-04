package leetcode

import "testing"

func TestAddStrings(t *testing.T) {
	cases := []struct {
		desc string
		num1 string
		num2 string
		want string
	}{
		{
			desc: "#1",
			num1: "123",
			num2: "456",
			want: "579",
		},
		{
			desc: "#2",
			num1: "777",
			num2: "888",
			want: "1665",
		},
		{
			desc: "#3",
			num1: "999",
			num2: "1",
			want: "1000",
		},
		{
			desc: "#4",
			num1: "5",
			num2: "0",
			want: "5",
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := addStrings(s.num1, s.num2); s.want != got {
				t.Errorf("num1: %s, num2: %s, want: %s, got: %s", s.num1, s.num2, s.want, got)
			}
		})
	}
}
