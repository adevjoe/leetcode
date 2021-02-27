package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFourSum(t *testing.T) {
	cases := []struct {
		arg1 []int
		arg2 int
		want [][]int
	}{
		{arg1: []int{1, 0, -1, 0, -2, 2}, arg2: 0, want: [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		{arg1: []int{-1, -5, -5, -3, 2, 5, 0, 4}, arg2: -7, want: [][]int{{-5, -5, -1, 4}, {-5, -3, -1, 2}}},
		{arg1: []int{0, 0, 0, 0}, arg2: 0, want: [][]int{{0, 0, 0, 0}}},
		{arg1: []int{}, arg2: 0, want: [][]int{}},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("TestFourSum-%d", i), func(t *testing.T) {
			if get := fourSum(c.arg1, c.arg2); !reflect.DeepEqual(get, c.want) {
				t.Errorf("arg1: %v, arg2: %v, want: %v, get: %v", c.arg1, c.arg2, c.want, get)
			}
		})
	}
}
