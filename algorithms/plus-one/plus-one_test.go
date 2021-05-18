package leetcode

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	cases := []struct {
		desc   string
		digits []int
		want   []int
	}{
		{
			desc:   "#1",
			digits: []int{1, 2, 3},
			want:   []int{1, 2, 4},
		},
		{
			desc:   "#2",
			digits: []int{4, 3, 2, 1},
			want:   []int{4, 3, 2, 2},
		},
		{
			desc:   "#3",
			digits: []int{0},
			want:   []int{1},
		},
		{
			desc:   "#4",
			digits: []int{9, 9, 9, 9},
			want:   []int{1, 0, 0, 0, 0},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := plusOne(s.digits); !reflect.DeepEqual(s.want, got) {
				t.Errorf("digits: %v, want: %v, got: %v", s.digits, s.want, got)
			}
		})
	}
}
