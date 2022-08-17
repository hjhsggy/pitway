package algorithm

func BubbleSort1(array []int) {

	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] < array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}

func InsertSort(array []int) {

	for i := 1; i < len(array); i++ {
		for j := i; j > 0; j-- {
			if array[j] > array[j-1] {
				array[j], array[j-1] = array[j-1], array[j]
			}
		}
	}

}

func SelectSort(array []int) {

	for i := 1; i < len(array); i++ {

	}

}

func QuickSort(array []int) []int {

	if len(array) == 0 {
		return array
	}

	low, mid, high := make([]int, 0), make([]int, 0), make([]int, 0)

	key := array[0]
	mid = append(mid, key)
	for i := 1; i < len(array); i++ {
		if array[i] == key {
			mid = append(mid, array[i])
		} else if array[i] < key {
			low = append(low, array[i])
		} else {
			high = append(high, array[i])
		}
	}

	low, high = QuickSort(low), QuickSort(high)

	return append(append(low, mid...), high...)

}

func QuickSort2(values []int) {

	length := len(values)

	if length <= 1 {
		return
	}

	mid, i := values[0], 1    // 取第一个元素作为分水岭，i下标初始为1，即分水岭右侧的第一个元素的下标
	head, tail := 0, length-1 // 头尾的下标

	// 如果头和尾没有相遇，就会一直触发交换
	for head < tail {
		if values[i] > mid {
			// 如果分水岭右侧的元素大于分水岭，就将右侧的尾部元素和分水岭右侧元素交换
			values[i], values[tail] = values[tail], values[i]
			tail-- // 尾下标左移一位
		} else {
			// 如果分水岭右侧的元素小于等于分水岭，就将分水岭右移一位
			values[i], values[head] = values[head], values[i]
			head++ // 头下标右移一位
			i++    // i下标右移一位
		}
	}

	// 分水岭左右的元素递归做同样处理
	QuickSort2(values[:head])
	QuickSort2(values[head+1:])

}

//堆排序
//s[0]不用，实际元素从角标1开始
//父节点元素大于子节点元素
//左子节点角标为2*k
//右子节点角标为2*k+1
//父节点角标为k/2
func HeapSort(s []int) {

	// 构造大根堆
	for i := len(s)/2 - 1; i >= 0; i-- {
		sink(s, i, len(s))
	}

	//调整堆顶与堆底元素, 构造大根堆
	for j := len(s) - 1; j >= 0; j-- {
		// 交换对顶与堆底元素
		swap(s, 0, j)
		// 继续让结构满足大根堆
		sink(s, 0, j)
	}

}

//下沉（由上至下的堆有序化）
func sink(s []int, i, N int) {

	tmp := s[i]

	for k := i*2 + 1; k < N; k = 2*k + 1 {
		if k+1 < N && s[k] < s[k+1] {
			k++
		}

		if s[k] > tmp {
			s[i] = s[k]
			i = k
		} else {
			break
		}
	}
	s[i] = tmp
}

func swap(s []int, i int, j int) {
	s[i], s[j] = s[j], s[i]
}
