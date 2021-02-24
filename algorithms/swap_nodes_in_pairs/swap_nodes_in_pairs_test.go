package swap_nodes_in_pairs

import (
	"reflect"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	cases := []struct {
		desc string
		head *ListNode
		want *ListNode
	}{
		{
			desc: "#1",
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}},
			want: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := swapPairs(s.head); !reflect.DeepEqual(s.want, got) {
				t.Error("error")
			}
		})
	}
}
