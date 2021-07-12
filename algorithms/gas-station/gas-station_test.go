package leetcode

import "testing"

func TestGasStation(t *testing.T) {
	cases := []struct {
		desc string
		gas  []int
		cost []int
		want int
	}{
		{
			desc: "#1",
			gas:  []int{1, 2, 3, 4, 5},
			cost: []int{3, 4, 5, 1, 2},
			want: 3,
		},
		{
			desc: "#2",
			gas:  []int{2, 3, 4},
			cost: []int{3, 4, 3},
			want: -1,
		},
		{
			desc: "#3",
			gas:  []int{3, 1, 1},
			cost: []int{1, 2, 2},
			want: 0,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := canCompleteCircuit(s.gas, s.cost); got != s.want {
				t.Error()
			}
		})
	}
}
