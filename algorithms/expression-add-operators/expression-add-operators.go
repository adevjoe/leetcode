// https://leetcode.com/problems/expression-add-operators

package leetcode

import "strconv"

var result []string

func addOperators(num string, target int) []string {
	if num == "" {
		return nil
	}
	result = nil
	addOperatorsSub(num, target, 0, 0, 0, "")
	return result
}

func addOperatorsSub(num string, target int, pre, preValue, pos int, path string) {
	if pos == len(num) {
		if pre == target {
			result = append(result, path)
		}
		return
	}
	for i := pos; i < len(num); i++ {
		if len(num[pos:i+1]) > 1 && num[pos] == '0' {
			continue
		}
		n, _ := strconv.Atoi(num[pos : i+1])
		if pos == 0 {
			addOperatorsSub(num, target, n, n, i+1, num[pos:i+1])
		} else {
			addOperatorsSub(num, target, pre+n, n, i+1, path+"+"+num[pos:i+1])
			addOperatorsSub(num, target, pre-n, -n, i+1, path+"-"+num[pos:i+1])
			addOperatorsSub(num, target, pre-preValue+preValue*n, preValue*n, i+1, path+"*"+num[pos:i+1])
		}
	}
}
