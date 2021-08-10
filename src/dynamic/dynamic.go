package dynamic

/*
给定一个字符串 (s) 和一个字符模式 (p)。实现支持 '.' 和 '*' 的正则表达式匹配。

'.' 匹配任意单个字符。

'*' 匹配零个或多个前面的元素。

匹配应该覆盖整个字符串 (s) ，而不是部分字符串。

说明:

s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。

*/

func MatchRegex(sub, pattern []string) bool {

	if len(sub) == 0 {
		return true
	}

	for i := 0; i < len(pattern); i++ {

		if pattern[i] == "*" {
			return true
		}

	}

	return false

}

/*
给定两个字符串s1和s2, 计算把s1转换为s2需要的最少操作次数, 可以对一个字符串进行插入,删除,替换一个字符
*/

func MinDistance(s1, s2 string) {

}
