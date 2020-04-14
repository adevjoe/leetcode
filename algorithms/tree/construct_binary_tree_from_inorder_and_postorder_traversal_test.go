package tree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBuildTreePostOrder(t *testing.T) {
	testCases := []struct {
		inorder   []int
		postorder []int
	}{
		{inorder: []int{9, 3, 15, 20, 7}, postorder: []int{9, 15, 7, 20, 3}},
		{inorder: []int{2, 3, 1}, postorder: []int{3, 2, 1}},
		{inorder: []int{3, 2, 1}, postorder: []int{3, 2, 1}},
		{inorder: []int{1, 2, 3}, postorder: []int{3, 2, 1}},
	}
	for i, c := range testCases {
		t.Run(fmt.Sprintf("buildTree%d", i), func(t *testing.T) {
			tree := buildTreePostOrder(c.inorder, c.postorder)
			if get := inorderTraversal(tree); !reflect.DeepEqual(c.inorder, get) {
				t.Errorf("build tree error, want %v, get %v", c.inorder, get)
			}
			if get := postorderTraversal(tree); !reflect.DeepEqual(c.postorder, get) {
				t.Errorf("build tree error, want %v, get %v", c.postorder, get)
			}
		})
	}

	for i, c := range testCases {
		t.Run(fmt.Sprintf("buildTree2-%d", i), func(t *testing.T) {
			tree := buildTreePostOrder2(c.inorder, c.postorder)
			if get := inorderTraversal(tree); !reflect.DeepEqual(c.inorder, get) {
				t.Errorf("build tree error, want %v, get %v", c.inorder, get)
			}
			if get := postorderTraversal(tree); !reflect.DeepEqual(c.postorder, get) {
				t.Errorf("build tree error, want %v, get %v", c.postorder, get)
			}
		})
	}
}
