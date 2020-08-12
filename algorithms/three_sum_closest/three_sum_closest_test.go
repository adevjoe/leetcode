package three_sum_closest

import (
	"fmt"
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	cases := []struct {
		arg1 []int
		arg2 int
		want int
	}{
		{arg1: []int{-1, 2, 1, -4}, arg2: 1, want: 2},
		{arg1: []int{50, 2, 1, -4}, arg2: 5, want: -1},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("TestThreeSumClosest-%d", i), func(t *testing.T) {
			if get := threeSumClosest(c.arg1, c.arg2); get != c.want {
				t.Errorf("arg1: %v, arg2: %v, want: %v, get: %v", c.arg1, c.arg2, c.want, get)
			}
		})
	}
}
