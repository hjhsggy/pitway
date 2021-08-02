package golang

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	arr := []int{1, 2, 3, 4}

	slice2 := arr[1:3]
	slice2[1] = 7

	slice2 = append(slice2, 10, 11)
	slice2[3] = 12

	fmt.Println(arr)
	fmt.Println(slice2)

}

func TestMap(t *testing.T) {

	strMap := make(map[string]string)
	fmt.Println(strMap)

	fmt.Println(1 << 3)

	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	rotate(nums, k)
	fmt.Println(nums)

}

func rotate(nums []int, k int) {

	k = k % len(nums)
	nums = append(nums[len(nums)-k:], nums[0:len(nums)-k]...)
}
