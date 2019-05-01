/*
Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.

Example 1:

Input: 121
Output: true
Example 2:

Input: -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
Follow up:

Coud you solve it without converting the integer to a string?
*/
package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(0))
	fmt.Println(isPalindrome(1))
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(11))
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(1122))
}

func isPalindrome(x int) bool {
	// filter
	// -121 is false
	if x < 0 {
		return false
	}
	// 10 110 4440 ... false
	if x%10 == 0 && x != 0 {
		return false
	}

	new := 0
	current := x

	// convert position backward as forward
	for new < x {
		new = new*10 + current%10
		current /= 10
	}

	if x != new {
		return false
	}
	return true
}
