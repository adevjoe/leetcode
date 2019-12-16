package palindrome_number

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(0))
	fmt.Println(isPalindrome(1))
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(11))
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(1122))
}
