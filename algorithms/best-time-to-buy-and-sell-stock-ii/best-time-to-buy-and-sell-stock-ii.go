// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii

package leetcode

func maxProfit(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}
