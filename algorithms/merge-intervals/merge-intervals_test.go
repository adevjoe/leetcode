package leetcode

import (
	"reflect"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	cases := []struct {
		desc      string
		intervals [][]int
		want      [][]int
	}{
		{
			desc:      "#1",
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			want:      [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			desc:      "#2",
			intervals: [][]int{{1, 4}, {4, 5}},
			want:      [][]int{{1, 5}},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := merge(s.intervals); !reflect.DeepEqual(got, s.want) {
				t.Errorf("intervals: %v, want: %v, got: %v", s.intervals, s.want, got)
			}
		})
	}
}
