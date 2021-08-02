package leetcode

import "testing"

func TestExcelSheetColumnNumber(t *testing.T) {
	cases := []struct {
		desc        string
		columnTitle string
		want        int
	}{
		{
			desc:        "#1",
			columnTitle: "A",
			want:        1,
		},
		{
			desc:        "#2",
			columnTitle: "AA",
			want:        27,
		},
		{
			desc:        "#3",
			columnTitle: "AB",
			want:        28,
		},
		{
			desc:        "#4",
			columnTitle: "ZY",
			want:        701,
		},
		{
			desc:        "#5",
			columnTitle: "FXSHRXW",
			want:        2147483647,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := titleToNumber(s.columnTitle); got != s.want {
				t.Error()
			}
		})
	}
}
