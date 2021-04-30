// https://leetcode.com/problems/group-anagrams

package leetcode

func groupAnagrams(strs []string) [][]string {
	// 26 个字母分别以 26 个质数表示，每个字符串数组的乘积相同则拥有相同的字母
	seeds := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 107}
	m := make(map[int][]int)
	for i := range strs {
		tmp := 1
		for j := range strs[i] {
			tmp *= seeds[int(strs[i][j] - 'a')]
		}
		m[tmp] = append(m[tmp], i)
	}
	result := make([][]string, 0)
	for i := range m {
		tmp := make([]string, 0)
		for j := range m[i] {
			tmp = append(tmp, strs[m[i][j]])
		}
		result = append(result, tmp)
	}
	return result
}
