package testutil

import "testing"

type Interface interface {
	Len() int
	Fn(i int)
	Args(i int) interface{}
	Want(i int) interface{}
	Get(i int) interface{}
	Assert(i int) bool
}

func Testing(t *testing.T, data Interface) {
	success := 0
	for i := 0; i < data.Len(); i++ {
		data.Fn(i)
		if data.Assert(i) {
			success++
		} else {
			t.Errorf("%s error, args: %+v, want: %+v, get: %+v", t.Name(), data.Args(i),
				data.Want(i), data.Get(i))
		}
	}
	t.Logf("%s success rate: %.2f%%", t.Name(), float64(success)/float64(data.Len())*100)
}
