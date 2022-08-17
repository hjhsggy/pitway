/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 *
 * https://leetcode-cn.com/problems/climbing-stairs/description/
 *
 * algorithms
 * Easy (53.60%)
 * Likes:    2535
 * Dislikes: 0
 * Total Accepted:    892.8K
 * Total Submissions: 1.7M
 * Testcase Example:  '2'
 *
 * 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
 *
 * 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 2
 * 输出：2
 * 解释：有两种方法可以爬到楼顶。
 * 1. 1 阶 + 1 阶
 * 2. 2 阶
 *
 * 示例 2：
 *
 *
 * 输入：n = 3
 * 输出：3
 * 解释：有三种方法可以爬到楼顶。
 * 1. 1 阶 + 1 阶 + 1 阶
 * 2. 1 阶 + 2 阶
 * 3. 2 阶 + 1 阶
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 45
 *
 *
 */
package leetcode

// 自底向上
/*
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	dp_1, dp_2 := 2, 1
	for i := 3; i <= n; i++ {
		dp_i := dp_1 + dp_2
		dp_2 = dp_1
		dp_1 = dp_i
	}
	return dp_1
}
*/

// 自顶向下
/*

 */

// @lc code=start
var memo []int

func climbStairs(n int) int {
	memo = make([]int, n+1)
	return dp(n)
}

func dp(n int) int {
	if n <= 2 {
		return n
	}
	if memo[n] > 0 {
		return memo[n]
	}
	memo[n] = dp(n-1) + dp(n-2)
	return memo[n]
}

// @lc code=end
