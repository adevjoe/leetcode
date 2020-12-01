// https://leetcode.com/problems/fruit-into-baskets/

package fruit_into_baskets

func totalFruit(tree []int) int {
	var i, p1, p2, total int
	for i < len(tree) {
		if tree[i] == tree[p1] || tree[i] == tree[p2] {
			if i-p1+1 > total {
				total = i - p1 + 1
			}
			i++
		} else {
			if p1 == p2 {
				if i-p1+1 > total {
					total = i - p1 + 1
				}
				p2 = i
				i++
			} else {
				i = p2
				p1 = p2
			}
		}
	}
	return total
}
