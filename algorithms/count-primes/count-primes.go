// https://leetcode.com/problems/count-primes

package leetcode

func countPrimes(n int) int {

	primeArr := make([]bool, n)

	for i := 2; i*i < n; i++ {
		if primeArr[i] == false {
			for j := i; j*i < n; j++ {
				primeArr[i*j] = true
			}
		}
	}
	count := 0
	if n > 2 {
		count++
	}
	for i := 3; i < n; i += 2 {
		if primeArr[i] == false {
			count++
		}
	}
	return count
}
