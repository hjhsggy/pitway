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

	low = QuickSort(low)
	high = QuickSort(high)

	return append(append(low, mid...), high...)

}
