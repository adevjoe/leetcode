package leetcode

import (
	"fmt"
	"testing"
)

var str = "   -422w"

func TestAtoi(t *testing.T) {

	fmt.Println(myAtoi(str))
	fmt.Println(atoi(str))
}

func BenchmarkAtoi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		atoi(str)
	}
}

func BenchmarkAtoi2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myAtoi(str)
	}
}
