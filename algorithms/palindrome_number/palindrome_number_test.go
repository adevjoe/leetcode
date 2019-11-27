package palindrome_number

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	fmt.Println(IsPalindrome(-121))
	fmt.Println(IsPalindrome(0))
	fmt.Println(IsPalindrome(1))
	fmt.Println(IsPalindrome(10))
	fmt.Println(IsPalindrome(11))
	fmt.Println(IsPalindrome(121))
	fmt.Println(IsPalindrome(1122))
}
