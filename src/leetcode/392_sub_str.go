package leetcode

import "fmt"

/*
给定字符串 s 和 t ，判断 s 是否为 t 的子序列。

字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。

进阶：

如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？

示例 1：

输入：s = "abc", t = "ahbgdc"
输出：true
示例 2：

输入：s = "axc", t = "ahbgdc"
输出：false
*/

func isSubsequence(s string, t string) bool {

	strMap := make(map[byte][]int)

	for i := range t {
		strMap[t[i]] = append(strMap[t[i]], i)
	}
	fmt.Println(strMap)

	for i := range s {
		// 不存在直接返回false
		if _, ok := strMap[s[i]]; !ok {
			return false
		}

		//

	}

	return true
}

func findKthLargest(nums []int, k int) int {

	for i := 0; i < k; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	return nums[k-1]

}
