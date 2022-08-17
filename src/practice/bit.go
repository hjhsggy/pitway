package practice

// 位运算

// 判断某个数是不是2的幂次方
func is2Power(n int64) bool {
	return n > 0 && n&(n-1) == 0
}
