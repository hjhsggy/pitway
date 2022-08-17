package leetcode

import (
	"sort"
	"strconv"
)

// 常规解法
func TwoSumBasic(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{-1, -1}
}

// 进阶
// 常规解法
func TwoSum(nums []int, target int) []int {

	return []int{-1, -1}
}

func RemoveCoveredIntervals(intervals [][]int) int {

	// 按照起点升序排列，起点相同时降序排列
	sort.Slice(intervals, func(i, j int) bool {
		// 起点相同, 结束点降序
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		// 起点升序
		return intervals[i][0] < intervals[j][0]
	})

	left := intervals[0][0]
	right := intervals[0][1]
	index := 0
	for _, val := range intervals {

		if left <= val[0] && right >= val[1] {
			index++
		}
		// 情况二，找到相交区间，合并
		if right >= val[0] && right <= val[1] {
			right = val[1]
		}
		// 情况三，完全不相交，更新起点和终点
		if right < val[0] {
			left = val[0]
			right = val[1]
		}

	}

	return len(intervals) - index + 1
}

func multiply(num1 string, num2 string) string {

	if num1 == "" || num2 == "" {
		return ""
	}

	l1, l2 := len(num1), len(num2)
	size := l1 + l2

	// 两数相乘最多为 l1 + l2
	result := make([]int, size)

	for i := l1 - 1; i >= 0; i-- {
		for j := l2 - 1; j >= 0; j-- {

			op1, _ := strconv.Atoi(string(num1[i]))
			op2, _ := strconv.Atoi(string(num2[j]))

			mut := op1 * op2
			index1 := size - 1 - (i + j)
			index2 := size - 1 - (i + j + 1)

			sum := mut + result[index2]

			result[index2] = sum % 10
			result[index1] += sum / 10

		}
	}

	var str string
	for i := size - 1; i >= 0; i-- {
		if i == size-1 && result[i] == 0 {
			continue
		}
		str += strconv.Itoa(result[i])
	}

	return str

}
