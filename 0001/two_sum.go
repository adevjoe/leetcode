/**
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/
package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for key, value := range nums {
		if (len(nums) - key) > 0 {
			for i := key + 1; i < len(nums); i++ {
				if (value + nums[i]) == target {
					return []int{key, i}
				}
			}
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
