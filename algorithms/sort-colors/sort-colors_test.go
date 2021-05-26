package leetcode

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	cases := []struct {
		desc string
		nums []int
		want []int
	}{
		{
			desc: "#1",
			nums: []int{2, 0, 2, 1, 1, 0},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			desc: "#2",
			nums: []int{0},
			want: []int{0},
		},
		{
			desc: "#3",
			nums: []int{2, 0, 1},
			want: []int{0, 1, 2},
		},
		{
			desc: "#4",
			nums: []int{1},
			want: []int{1},
		},
		{
			desc: "#5",
			nums: []int{1, 0},
			want: []int{0, 1},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			origin := make([]int, len(s.nums))
			copy(origin, s.nums)
			if sortColors(s.nums); !reflect.DeepEqual(s.nums, s.want) {
				t.Errorf("nums: %v, want: %v, got: %v", origin, s.want, s.nums)
			}
		})
	}
}
