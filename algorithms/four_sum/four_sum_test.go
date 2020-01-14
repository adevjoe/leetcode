package four_sum

import (
	"github.com/adevjoe/leetcode/testutil"
	"reflect"
	"testing"
)

type TestCase struct {
	Args struct {
		nums   []int
		target int
	}
	Want [][]int
	Get  [][]int
}

type TestCaseSlice []TestCase

func (p TestCaseSlice) Len() int               { return len(p) }
func (p TestCaseSlice) Fn(i int)               { p[i].Get = fourSum(p[i].Args.nums, p[i].Args.target) }
func (p TestCaseSlice) Args(i int) interface{} { return p[i].Args }
func (p TestCaseSlice) Want(i int) interface{} { return p[i].Want }
func (p TestCaseSlice) Get(i int) interface{}  { return p[i].Get }
func (p TestCaseSlice) Assert(i int) bool {
	return reflect.DeepEqual(p[i].Get, p[i].Want)
}

func TestFourSum(t *testing.T) {
	data := TestCaseSlice{
		{Args: struct {
			nums   []int
			target int
		}{nums: []int{1, 0, -1, 0, -2, 2}, target: 0}, Want: [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		{Args: struct {
			nums   []int
			target int
		}{nums: []int{-1, -5, -5, -3, 2, 5, 0, 4}, target: -7}, Want: [][]int{{-5, -5, -1, 4}, {-5, -3, -1, 2}}},
		{Args: struct {
			nums   []int
			target int
		}{nums: []int{0, 0, 0, 0}, target: 0}, Want: [][]int{{0, 0, 0, 0}}},
		{Args: struct {
			nums   []int
			target int
		}{nums: []int{}, target: 0}, Want: [][]int{}},
	}
	testutil.Testing(t, data)
}
