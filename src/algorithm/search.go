package algorithm

func BinarySearch(array []int, target int) int {

	if len(array) == 0 {
		return -1
	}

	left := 0
	right := len(array) - 1

	for i := 0; left <= right; i++ {
		mid := left + (right-left)/2

		if array[mid] == target {
			return mid
		}

		if array[mid] < target {
			left = mid + 1
		}

		if array[mid] > target {
			right = mid - 1
		}

	}

	return -1

}
