// https://leetcode.com/problems/serialize-and-deserialize-binary-tree/

package tree

import (
	"strconv"
	"strings"
)

type CodecSD struct {
	cur int
	sb  strings.Builder
}

func ConstructorSD() CodecSD {
	return CodecSD{sb: strings.Builder{}}
}

// Serializes a tree to a single string.
func (this *CodecSD) serialize(root *TreeNode) string {
	this.buildString(root)
	return this.sb.String()
}

func (this *CodecSD) buildString(root *TreeNode) {
	if root == nil {
		this.sb.Write([]byte(","))
		return
	}
	this.sb.Write([]byte(strconv.Itoa(root.Val)))
	this.sb.Write([]byte(","))
	this.buildString(root.Left)
	this.buildString(root.Right)
}

// Deserializes your encoded data to tree.
func (this *CodecSD) deserialize(data string) *TreeNode {
	if data[this.cur] == ',' {
		this.cur++
		return nil
	}
	if this.cur >= len(data) {
		return nil
	}
	tmp := ""
	for data[this.cur] != ',' {
		tmp += string(data[this.cur])
		this.cur++
	}
	val, _ := strconv.Atoi(tmp)
	t := &TreeNode{Val: val}
	this.cur++
	t.Left = this.deserialize(data)
	t.Right = this.deserialize(data)
	return t
}

func (this *CodecSD) deserialize2(data string) *TreeNode {
	if data[this.cur] == ',' {
		this.cur++
		return nil
	}
	if this.cur >= len(data) {
		return nil
	}
	tmp := ""
	for data[this.cur] != ',' {
		tmp += string(data[this.cur])
		this.cur++
	}
	val, _ := strconv.Atoi(tmp)
	t := &TreeNode{Val: val}
	this.cur++
	t.Left = this.deserialize2(data)
	t.Right = this.deserialize2(data)
	return t
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
