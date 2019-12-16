package int_to_roman

import (
	"fmt"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	fmt.Println(intToRoman(3))
	fmt.Println(intToRoman(4))
	fmt.Println(intToRoman(9))
	fmt.Println(intToRoman(40))
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(1994))

	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("XL"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}
