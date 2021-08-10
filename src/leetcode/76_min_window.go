package leetcode

// minWindow 滑动窗口
func minWindow(s, t string) string {

	// 目标串和模式串任意一个为空直接返回
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return ""
	}

	need, window := make(map[byte]uint8), make(map[byte]uint8)
	for i := range t {
		need[t[i]]++
	}

	var left, right, valid, start, min int = 0, 0, 0, 0, len(s)
	for right < len(s) {
		tempAdd := s[right]
		right++

		if _, ok := need[tempAdd]; ok {
			window[tempAdd]++
			if window[tempAdd] == need[tempAdd] {
				valid++
			}
		}

		for valid == len(need) {
			// min 记录满足要求的最短字串
			if right-left < min {
				start = left
				min = right - left
			}

			tempDel := s[left]
			left++

			if _, ok := need[tempDel]; ok {
				if window[tempDel] == need[tempDel] {
					valid--
				}
				window[tempDel]--
			}
		}
	}
	return string(s[start : start+min])

}

func checkInclus1ion(s1 string, s2 string) bool {

	need, window := make(map[byte]uint8), make(map[byte]uint8)
	for i := range s1 {
		need[s1[i]]++
	}

	left, right := 0, 0
	valid := 0
	for right < len(s2) {
		tempAdd := s2[right]
		right++

		if _, ok := need[tempAdd]; ok {
			window[tempAdd]++
			if window[tempAdd] == need[tempAdd] {
				valid++
			}
		}

		for right-left >= len(s1) {
			if valid == len(need) {
				return true
			}
			tempDel := s2[left]
			left++

			if _, ok := need[tempDel]; ok {
				if window[tempDel] == need[tempDel] {
					valid--
				}
				window[tempDel]--
			}
		}
	}
	return false

}

func findAnagrams(s string, p string) []int {

	need, window := map[byte]int{}, map[byte]int{}
	for i := range p {
		need[p[i]]++
	}

	left, right := 0, 0
	valid := 0
	indexVector := []int{}
	for right < len(s) {
		tmpAdd := s[right]
		right++
		if _, ok := need[tmpAdd]; ok {
			window[tmpAdd]++
			if window[tmpAdd] == need[tmpAdd] {
				valid++
			}
		}
		for right-left >= len(p) {
			if valid == len(need) {
				indexVector = append(indexVector, left)
			}
			tmpDel := s[left]
			left++
			if _, ok := window[tmpDel]; ok {
				if window[tmpDel] == need[tmpDel] {
					valid--
				}
				window[tmpDel]--
			}
		}
	}
	return indexVector
}

func findAnagram(s string, p string) []int {

	need, window := map[byte]int{}, map[byte]int{}
	for i := range p {
		need[p[i]]++
	}
	left, right := 0, 0
	valid := 0
	indexs := []int{}
	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		for right-left >= len(p) {
			if valid == len(need) {
				indexs = append(indexs, left)
			}
			d := s[left]
			left++
			if _, ok := window[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return indexs

}
