package string_add

import (
	"fmt"
	"testing"
)

func TestStringAdd(t *testing.T) {
	cases := []struct {
		arg1 string
		arg2 string
		want string
	}{
		{
			arg1: "001", arg2: "1", want: "",
		},
		{
			arg1: "1", arg2: "x", want: "",
		},
		{
			arg1: "123", arg2: "456", want: "579",
		},
		{
			arg1: "999", arg2: "1", want: "1000",
		},
		{
			arg1: "109", arg2: "1", want: "110",
		},
		{
			arg1: "555", arg2: "555", want: "1110",
		},
		{
			arg1: "55", arg2: "555", want: "610",
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("TestStringAdd-%d", i), func(t *testing.T) {
			if get := stringAdd(c.arg1, c.arg2); get != c.want {
				t.Errorf("arg1: %v, arg2: %v, get: %v, want: %v", c.arg1, c.arg2, get, c.want)
			}
		})
	}
}

func BenchmarkStringAdd(b *testing.B) {
	num1 := "555"
	num2 := "555"
	b.Run("stringAdd", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			stringAdd(num1, num2)
		}
	})
	b.Run("stringAddForByte", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			stringAddForByte(num1, num2)
		}
	})
}
