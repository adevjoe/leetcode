// https://leetcode.com/problems/spiral-matrix

package leetcode

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 1 {
		return matrix[0]
	}
	m := len(matrix)
	n := len(matrix[0])
	i := 0
	j := 0
	result := make([]int, 0)
	cicle := 1
	l := m
	if n < m {
		l = n
	}
	maxCicle := (l + 1) / 2
	for cicle <= maxCicle {
		// right
		for j <= n-cicle {
			result = append(result, matrix[i][j])
			j++
			if j > n-cicle {
				j--
				break
			}
		}
		i++
		if i > m-cicle {
			break
		}

		// down
		for i <= m-cicle {
			result = append(result, matrix[i][j])
			i++
			if i > m-cicle {
				i--
				break
			}
		}
		j--
		if j < cicle-1 {
			break
		}

		// left
		for j >= cicle-1 {
			result = append(result, matrix[i][j])
			j--
			if j < cicle-1 {
				j++
				break
			}
		}
		i--

		// up
		for i > cicle-1 {
			result = append(result, matrix[i][j])
			i--
			if i == cicle-1 {
				i++
				break
			}
		}
		j++
		cicle++
	}
	return result
}
