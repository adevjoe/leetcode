package testutil

import (
	"testing"
)

func GetSelf(i int) int {
	return i
}

type TestCase struct {
	Args int
	Want int
	Get  int
}

type TestCaseSlice []TestCase

func (p TestCaseSlice) Len() int               { return len(p) }
func (p TestCaseSlice) Fn(i int)               { p[i].Get = GetSelf(p[i].Args) }
func (p TestCaseSlice) Args(i int) interface{} { return p[i].Args }
func (p TestCaseSlice) Want(i int) interface{} { return p[i].Want }
func (p TestCaseSlice) Get(i int) interface{}  { return p[i].Get }
func (p TestCaseSlice) Assert(i int) bool      { return p[i].Get == p[i].Want }

func TestGetSelf(t *testing.T) {
	data := TestCaseSlice{{Args: 1, Want: 1}, {Args: 2, Want: 2}}
	Testing(t, data)
}
