package string_to_integer

import (
	"fmt"
	"testing"
)

var str = "   -422w"

func TestAtoi(t *testing.T) {

	fmt.Println(myAtoi(str))
	fmt.Println(Atoi(str))
}

func BenchmarkAtoi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Atoi(str)
	}
}

func BenchmarkAtoi2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myAtoi(str)
	}
}
