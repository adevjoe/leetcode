// https://leetcode.com/problems/expression-add-operators/

package algorithms

import "strconv"

var result []string

func addOperators(num string, target int) []string {
	if num == "" {
		return nil
	}
	result = make([]string, 0)
	addOperatorsSub(num, target, 0, 0, "")
	return result
}

func addOperatorsSub(num string, target int, pre, pos int, path string) {
	if pos == len(num) {
		if pre == target {
			result = append(result, path)
		}
		return
	}
	for i := pos; i < len(num); i++ {
		n, _ := strconv.Atoi(num[pos : i+1])
		if pos == 0 {
			addOperatorsSub(num, target, n, i+1, num[pos:i+1])
		} else {
			addOperatorsSub(num, target, pre+n, i+1, path+"+"+num[pos:i+1])
			addOperatorsSub(num, target, pre-n, i+1, path+"-"+num[pos:i+1])
			addOperatorsSub(num, target, pre*n, i+1, path+"*"+num[pos:i+1])
		}
	}
}
