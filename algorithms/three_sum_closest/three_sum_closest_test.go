package three_sum_closest

import (
	"github.com/adevjoe/leetcode/testutil"
	"testing"
)

var testCase = []testutil.TestCase{
	{
		Args: []interface{}{
			[]int{-1, 2, 1, -4},
			1,
		}, Want: 2,
	},
	{
		Args: []interface{}{
			[]int{50, 2, 1, -4},
			5,
		}, Want: -1,
	},
}

func TestThreeSumClosest(t *testing.T) {
	for _, c := range testCase {
		if get := threeSumClosest(c.Args.([]interface{})[0].([]int), c.Args.([]interface{})[1].(int)); get != c.Want {
			t.Errorf("ThreeSumClosest error with nums: %v, target: %d, get: %v, want: %v",
				c.Args.([]interface{})[0].([]int), c.Args.([]interface{})[1].(int), get, c.Want)
		}
	}
}
