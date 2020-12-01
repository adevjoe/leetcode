package fruit_into_baskets

import "testing"

func TestTotalFruit(t *testing.T) {
	cases := []struct {
		desc string
		tree []int
		want int
	}{
		{
			desc: "#1",
			tree: []int{},
			want: 0,
		},
		{
			desc: "#2",
			tree: []int{1, 1},
			want: 2,
		},
		{
			desc: "#3",
			tree: []int{1, 2},
			want: 2,
		},
		{
			desc: "#4",
			tree: []int{1, 2, 1},
			want: 3,
		},
		{
			desc: "#5",
			tree: []int{0, 1, 2, 2},
			want: 3,
		},
		{
			desc: "#6",
			tree: []int{1, 2, 3, 2, 2},
			want: 4,
		},
		{
			desc: "#7",
			tree: []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4},
			want: 5,
		},
		{
			desc: "#8",
			tree: []int{0, 0, 1, 1},
			want: 4,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := totalFruit(s.tree); s.want != got {
				t.Errorf("tree: %v, want: %d, got: %d", s.tree, s.want, got)
			}
		})
	}
}
