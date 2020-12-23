package algorithms

import (
	"reflect"
	"testing"
)

func TestAddOperators(t *testing.T) {
	cases := []struct {
		desc   string
		num    string
		target int
		want   []string
	}{
		{
			desc:   "123",
			num:    "123",
			target: 6,
			want:   []string{"1+2+3", "1*2*3"},
		},
		{
			desc:   "232",
			num:    "232",
			target: 8,
			want:   []string{"2+3*2", "2*3+2"},
		},
		{
			desc:   "105",
			num:    "105",
			target: 5,
			want:   []string{"1*0+5", "10-5"},
		},
		{
			desc:   "00",
			num:    "00",
			target: 0,
			want:   []string{"0+0", "0-0", "0*0"},
		},
		{
			desc:   "not operator",
			num:    "123",
			target: 123,
			want:   []string{"123"},
		},
		{
			desc:   "3456237490",
			num:    "3456237490",
			target: 9191,
			want:   nil,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := addOperators(s.num, s.target); !reflect.DeepEqual(got, s.want) {
				t.Errorf("num: %s, target: %d, want: %v, got: %v", s.num, s.target, s.want, got)
			}
		})
	}
}
