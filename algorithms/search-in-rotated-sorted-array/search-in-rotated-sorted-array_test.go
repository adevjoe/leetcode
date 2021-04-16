package leetcode

import "testing"

func TestSearchinRotatedSortedArray(t *testing.T) {
	cases := []struct {
		desc   string
		nums   []int
		target int
		want   int
	}{
		{
			desc:   "simple",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			want:   4,
		},
		{
			desc:   "not found",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			want:   -1,
		},
		{
			desc:   "another not found",
			nums:   []int{1},
			target: 3,
			want:   -1,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := search(s.nums, s.target); got != s.want {
				t.Errorf("nums: %v, target: %d, want: %d, got: %d", s.nums, s.target, s.want, got)
			}
			if got := search1(s.nums, s.target); got != s.want {
				t.Errorf("nums: %v, target: %d, want: %d, got: %d", s.nums, s.target, s.want, got)
			}
		})
	}
}

func getNums(n int, max int) []int {
	var a []int
	for i := n; i <= max; i++ {
		a = append(a, i)
	}
	for i := 0; i < n; i++ {
		a = append(a, i)
	}
	return a
}

func BenchmarkSearch(b *testing.B) {
	nums := getNums(21321310, 21321311)

	for i := 0; i < b.N; i++ {
		search(nums, 21321309)
	}
}

func BenchmarkSearch1(b *testing.B) {
	nums := getNums(21321310, 21321311)

	for i := 0; i < b.N; i++ {
		search1(nums, 21321309)
	}
}

func TestTest(t *testing.T) {
	a := []int{1}
	t.Log(a[:1])
	t.Log(a[1:])
	b := []int{1, 2}
	t.Log(b[:2])
	t.Log(b[2:])
	c := []int{1, 2, 3}
	t.Log(c[:2])
	t.Log(c[2:])
	t.Log(3 / 2)
}
