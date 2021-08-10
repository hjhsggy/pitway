package leetcode

import (
	"sort"
	"testing"
)

func TestTwoSum(t *testing.T) {

	var nums = []int{2, 7, 12, 13}
	var target = 9

	result := TwoSumBasic(nums, target)

	t.Log(result)

}

type IntervalST struct {
	Start int
	End   int
}

// 按照 Person.Age 从大到小排序
type Interval []IntervalST

func (a Interval) Len() int { // 重写 Len() 方法
	return len(a)
}
func (a Interval) Swap(i, j int) { // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a Interval) Less(i, j int) bool { // 重写 Less() 方法, 从小到大排序

	// 起点相同, 结束点降序
	if a[i].Start == a[j].Start {
		return a[i].End > a[j].End
	}

	// 起点升序
	return a[i].Start < a[j].Start
}

func TestRemove(t *testing.T) {

	intervals := [][]int{
		{1, 4},
		{2, 6},
		{2, 8},
	}

	sort.Slice(intervals, func(i, j int) bool {
		// 起点相同, 结束点降序
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		// 起点升序
		return intervals[i][0] < intervals[j][0]
	})

	result := RemoveCoveredIntervals(intervals)

	t.Log(result)
}

func TestArraySorter(t *testing.T) {

	type IntervalST struct {
		Start int
		End   int
	}

	intervals := make([]*IntervalST, 0)

	data := [][]int{
		{1, 4},
		{2, 6},
		{2, 8},
	}

	for _, val := range data {
		intervals = append(intervals, &IntervalST{
			Start: val[0],
			End:   val[1],
		})
		t.Log(&intervals)
	}

	sort.Slice(intervals, func(i, j int) bool {
		// 起点相同, 结束点降序
		if intervals[i].Start == intervals[j].Start {
			return intervals[i].End > intervals[j].End
		}
		// 起点升序
		return intervals[i].Start < intervals[j].Start
	})

	for _, val := range intervals {
		t.Logf("start:%v -- end:%v", val.Start, val.End)
	}

}

func TestMinWindows(t *testing.T) {

	pat := "adobecodebanc"
	src := "abc"

	t.Log(minWindow(pat, src))

}

func TestCheckInclus1ion(t *testing.T) {

	// s1 := "ab"
	// s2 := "eidbvaooo"

	s1 := "abcdxabcde" // "ab"
	s2 := "abcdeabcdx" //"eidbaooo"

	t.Log(checkInclus1ion(s1, s2))

}

func TestFindAnagrams(t *testing.T) {

	// s1 := "ab"
	// s2 := "eidbvaooo"

	s1 := "abcdxabcde" // "ab"
	s2 := "abcd"       //"eidbaooo"

	t.Log(findAnagram(s1, s2))

	t.Log(findAnagrams(s1, s2))

}

func TestMutil(t *testing.T) {

	num1 := "9"
	num2 := "99"

	t.Log(multiply(num1, num2))

}

func Test392Substr(t *testing.T) {

	s, p := "abc", "adxbedc"

	isSubsequence(s, p)

}

func Test215KLargest(t *testing.T) {

	arr := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4

	t.Log(findKthLargest(arr, k))

}
