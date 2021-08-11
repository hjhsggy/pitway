package backtrack

/*
* 给定一个数列, 求其所有的全排列组合
 */

/*
 框架

result = []
def backtrack(路径, 选择列表):
    if 满足结束条件:
        result.add(路径)
        return

    for 选择 in 选择列表:
        做选择
        backtrack(路径, 选择列表)
        撤销选择
*/

func permute(nums []int) [][]int {

	// 初识化全局变量
	res := [][]int{}
	track := []int{}

	var backtrack func(nums []int, track []int)
	backtrack = func(nums []int, track []int) {
		if len(nums) == len(track) {
			temp := make([]int, len(track))
			// 为什么加入解集时，要将数组（描述中数组都是切片）内容复制到一个新的数组里，再加入解集?
			// 这个 track 变量是一个地址引用，结束当前递归，将它加入 res，后续的递归分支还要继续进行搜
			// 索，还要继续传递这个 track ，这个地址引用所指向的内存空间还要继续被操作，所以 res 中的
			// track 所引用的内容会被改变，这就造成了 res 中的内容随 track 变化。
			// 所以要复制 track 内容到一个新的数组里，然后放入 res，这样后续对 track 的操作，
			// 就不会影响已经放入 res 的内容。
			copy(temp, track)
			res = append(res, temp)
			return
		}
		for _, num := range nums {

			// 实现 contains() 来排除不合法的选择
			flag := false
			for _, a := range track {
				if a == num {
					flag = true
				}
			}
			if flag {
				continue
			}

			track = append(track, num)
			backtrack(nums, track)
			track = track[:len(track)-1]
		}
	}

	backtrack(nums, track)
	return res
}
