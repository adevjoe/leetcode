package int_to_roman

import (
	"fmt"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	fmt.Println(IntToRoman(3))
	fmt.Println(IntToRoman(4))
	fmt.Println(IntToRoman(9))
	fmt.Println(IntToRoman(40))
	fmt.Println(IntToRoman(58))
	fmt.Println(IntToRoman(1994))

	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("XL"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}
