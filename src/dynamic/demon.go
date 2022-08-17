package dynamic

import (
	"fmt"
)

func Fib(num int) int {

	if num == 1 || num == 2 {
		return 1
	}

	return Fib(num-1) + Fib(num-2)

}

func FibWithMemo(num int) int {

	if num < 1 {
		return 0
	}

	memo := make([]int, num+1)
	return helper(memo, num)
}

func helper(memo []int, num int) int {

	if (num == 1) || (num == 2) {
		return 1
	}

	// 中间数组保存记录
	if memo[num] != 0 {
		return memo[num]
	}

	memo[num] = helper(memo, num-1) + helper(memo, num-2)
	return memo[num]

}

func FibWithDP(num int) int {

	if num < 0 {
		return 0
	}
	if num == 1 || num == 2 {
		return 1
	}

	dp := make([]int, num+1)
	dp[1], dp[2] = 1, 1

	for i := 3; i <= num; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[num]

}

/*
* 给你 k 种面值的硬币，面值分别为 c1, c2 ... ck，每种硬币的数量无限，再给一个总金额 amount，
* 问你最少需要几枚硬币凑出这个金额，如果不可能凑出，算法返回 -1 。算法的函数签名如下：
* coins 中是可选硬币面值，amount 是目标金额
 */

var coins = []int{5, 2, 1}
var memo = make(map[int]int, len(coins))

func CoinChange(amount int) int {

	return Coin(amount)

}

func Coin(amount int) int {

	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}

	if _, ok := memo[amount]; ok {
		return memo[amount]
	}

	min := make([]int, 0)
	for _, coin := range coins {
		subProblem := Coin(amount - coin)
		if subProblem == -1 {
			continue
		}
		min = append(min, subProblem+1)
	}

	memo[amount] = minAmount(min)

	return memo[amount]
}

func minAmount(arr []int) int {

	minValue := arr[0]

	for _, val := range arr {

		if val < minValue {
			minValue = val
		}
	}

	return minValue

}

func CoinNums(amount int) {

	for _, coin := range coins {

		tmp := amount / coin
		amount = amount % coin

		fmt.Printf("coin:%v, num:%v\n", coin, tmp)
	}

}

/*
* ************************************************************************/
// 暴力解决最长公共子序列
func CommSubStr(txt, pat string) int {

	if len(txt) == 0 {
		return 0
	}

	i, j := 0, 0
	for i < len(txt) && j < len(pat) {
		if j == -1 || txt[i] == pat[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}

	}

	if j == len(pat) {
		return i - j
	} else {
		return -1
	}

}

// 求解相同前缀字串
func substringRepetition(needle string) []int {

	next := make([]int, len(needle))
	for i, j := 1, 0; i < len(needle); i++ {
		for j > 0 && needle[i] != needle[j] {
			// 上一个公共子序列
			j = next[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		next[i] = j
	}
	return next
}

func KMP(s string) [][256]int {

	db := make([][256]int, len(s))

	db[0][0] = 1

	x := 0

	for j := range s {
		for c := 0; c < 256; c++ {
			if int(s[j]) == c {
				db[j][c] = j + 1
			} else {
				db[j][c] = db[x][c]
			}
			x = db[x][s[j]]
		}
	}

	return db
}

func searchKMP(s, p string) int {

	db := KMP(s)

	j := 0
	for i := range p {
		j = db[j][p[i]]
		if j == len(s) {
			return i - len(s) + 1
		}
	}
	return -1
}
