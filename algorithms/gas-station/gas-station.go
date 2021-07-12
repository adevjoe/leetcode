// https://leetcode.com/problems/gas-station

package leetcode

func canCompleteCircuit(gas []int, cost []int) int {
	total := 0
	for i := 0; i < len(gas); i++ {
		total += gas[i] - cost[i]
	}
	if total < 0 {
		return -1
	}
	start := 0
	tank := 0
	for i := 0; i < len(gas); i++ {
		tank += gas[i] - cost[i]
		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}
	return start
}
