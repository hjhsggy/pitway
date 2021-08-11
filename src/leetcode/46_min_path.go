package leetcode

/*
矩阵从左上角到右下角的最短路径
*/

const MaxDistance = 2 ^ 32

func MinDistance(array [][]int) int {

	m := len(array)
	n := len(array[0])

	return distanceDP(array, m-1, n-1)

}

func distanceDP(array [][]int, i, j int) int {

	// base case
	if i == 0 && j == 0 {
		return array[0][0]
	}

	if i < 0 || j < 0 {
		return MaxDistance
	}

	left, top := distanceDP(array, i-1, j), distanceDP(array, i, j-1)
	if left < top {
		return left + array[i][j]
	} else {
		return top + array[i][j]
	}

}
