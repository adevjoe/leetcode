package three_sum_closest

import (
	"github.com/adevjoe/leetcode/testutil"
	"testing"
)

type ThreeSumClosestTestCase struct {
	Args struct {
		Nums   []int
		target int
	}
	Want int
	Get  int
}

type s []ThreeSumClosestTestCase

func (p s) Len() int               { return len(p) }
func (p s) Fn(i int)               { p[i].Get = threeSumClosest(p[i].Args.Nums, p[i].Args.target) }
func (p s) Args(i int) interface{} { return p[i].Args }
func (p s) Want(i int) interface{} { return p[i].Want }
func (p s) Get(i int) interface{}  { return p[i].Get }
func (p s) Assert(i int) bool      { return p[i].Get == p[i].Want }

func TestThreeSumClosest(t *testing.T) {
	data := s{
		{Args: struct {
			Nums   []int
			target int
		}{Nums: []int{-1, 2, 1, -4}, target: 1}, Want: 2},
		{Args: struct {
			Nums   []int
			target int
		}{Nums: []int{50, 2, 1, -4}, target: 5}, Want: -1},
	}
	testutil.Testing(t, data)
}
