package house_robber

import "testing"

func TestRobTwo(t *testing.T) {
	cases := []struct {
		desc string
		nums []int
		want int
	}{
		{
			desc: "sample",
			nums: []int{2, 3, 2},
			want: 3,
		},
		{
			desc: "sample #2",
			nums: []int{1, 2, 3, 1},
			want: 4,
		},
		{
			desc: "long",
			nums: []int{226, 174, 214, 16, 218, 48, 153, 131, 128, 17, 157, 142, 88, 43, 37, 157, 43, 221, 191, 68, 206, 23, 225, 82, 54, 118, 111, 46, 80, 49, 245, 63, 25, 194, 72, 80, 143, 55, 209, 18, 55, 122, 65, 66, 177, 101, 63, 201, 172, 130, 103, 225, 142, 46, 86, 185, 62, 138, 212, 192, 125, 77, 223, 188, 99, 228, 90, 25, 193, 211, 84, 239, 119, 234, 85, 83, 123, 120, 131, 203, 219, 10, 82, 35, 120, 180, 249, 106, 37, 169, 225, 54, 103, 55, 166, 124},
			want: 7102,
		},
		{
			desc: "null",
			nums: []int{},
			want: 0,
		},
		{
			desc: "zero",
			nums: []int{0},
			want: 0,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := robTwo(s.nums); got != s.want {
				t.Errorf("nums: %+v, want: %d, got: %d", s.nums, s.want, got)
			}
		})
	}
}
