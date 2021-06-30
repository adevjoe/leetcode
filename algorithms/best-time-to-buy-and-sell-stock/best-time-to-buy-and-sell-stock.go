// https://leetcode.com/problems/best-time-to-buy-and-sell-stock

package leetcode

var cache map[int]int

func maxProfit(prices []int) int {
	cache = make(map[int]int)
	min := prices[0]
	profit := 0
	for i := 1; i < len(prices); i++ {
		if a := prices[i] - min; a > profit {
			profit = a
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return profit
}
