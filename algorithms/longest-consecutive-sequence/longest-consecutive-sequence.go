// https://leetcode.com/problems/longest-consecutive-sequence

package leetcode

import "math"

func longestConsecutive(nums []int) int {
	head := make(map[int]int)
	tail := make(map[int]int)
	length := 0
	for _, n := range nums {
		if _, ok := head[n]; ok {
			continue
		}
		if _, ok := tail[n]; ok {
			continue
		}
		_, ok := head[n+1]
		_, ok2 := tail[n-1]
		if ok && ok2 {
			head[tail[n-1]] = head[n+1]
			tail[head[n+1]] = tail[n-1]
			if head[n+1]-tail[head[n+1]]+1 > length {
				length = head[n+1] - tail[head[n+1]] + 1
			}
			delete(head, n+1)
			delete(tail, n-1)
			continue
		}
		if _, ok := head[n+1]; ok {
			head[n] = head[n+1]
			tail[head[n+1]] = n
			if head[n+1]-tail[head[n+1]]+1 > length {
				length = head[n+1] - tail[head[n+1]] + 1
			}
			delete(head, n+1)
			continue
		}
		if _, ok := tail[n-1]; ok {
			head[tail[n-1]] = n
			tail[n] = tail[n-1]
			if n-tail[n]+1 > length {
				length = n - tail[n] + 1
			}
			delete(tail, n-1)
			continue
		}
		head[n] = n
		tail[n] = n
		if length == 0 {
			length = 1
		}
	}
	return length
}

// https://leetcode.com/problems/longest-consecutive-sequence/discuss/41057/Simple-O(n)-with-Explanation-Just-walk-each-streak
func longestConsecutive2(nums []int) int {
	numsMap := make(map[int]bool)
	best := 0
	for _, num := range nums {
		numsMap[num] = true
	}
	for num := range numsMap {
		if !numsMap[num-1] {
			y := num + 1
			for numsMap[y] {
				y++
			}
			best = int(math.Max(float64(best), float64(y-num)))
		}
	}
	return best
}
