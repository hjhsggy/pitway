package algorithm

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {

	array := []int{1, 3, 5, 7, 8, 9}
	target := 10

	index := BinarySearch(array, target)

	t.Log(index)

}

func TestSort(t *testing.T) {

	array := []int{5, 1, 3, 7, 4, 8, 2}

	// //test1 := array
	// test2 := array
	// //BubbleSort1(test1)
	// InsertSort(test2)

	// t.Log(QuickSort(array))

	//t.Log("bubble", test1, "\n")
	//t.Log("bubble", test2, "\n")

	QuickSort2(array)

	t.Log(array)

}

func TestLink(t *testing.T) {

	var head LinkList
	head.Insert(2)
	head.Insert(3)
	head.Insert(4)
	head.Insert(5)
	head.Insert(5)
	head.Insert(5)

	for p := head.Next; p != nil; p = p.Next {
		t.Log(p.Value)
	}

}

func TestLinkReverse(t *testing.T) {

	var head LinkList
	head.Insert("A")
	head.Insert("E")
	head.Insert("D")
	head.Insert("C")
	head.Insert("B")
	head.Insert("F")

	// for p := head.Next; p != nil; p = p.Next {
	// 	t.Log(p.Value)
	// }

	// result := head.Next.ReverseAll()
	// result.Print()

	// result := head.Next.ReversePreN(5)
	// result.Print()

	// result := head.Next.ReverseInterval(4, 5)
	// result.Print()

	// result = RemoveBackN(head.Next, 1)
	// result.Print()

	result := ListReverse(head.Next)
	result.Print()

}
