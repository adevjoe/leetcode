package leetcode

import "testing"

func TestEvaluateReversePolishNotation(t *testing.T) {
	cases := []struct {
		desc   string
		tokens []string
		want   int
	}{
		{
			desc:   "#1",
			tokens: []string{"2", "1", "+", "3", "*"},
			want:   9,
		},
		{
			desc:   "#2",
			tokens: []string{"4", "13", "5", "/", "+"},
			want:   6,
		},
		{
			desc:   "#3",
			tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			want:   22,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := evalRPN(s.tokens); got != s.want {
				t.Error()
			}
		})
	}
}
