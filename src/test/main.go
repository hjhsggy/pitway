package main

import "fmt"

type LinkList struct {
	Val  int
	Next *LinkList
}

func LinkReverse(head *LinkList) *LinkList {

	var pre *LinkList = nil
	var cur, next = head, head

	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

func (node *LinkList) headInsert(data int) {

	if node == nil {
		node = &LinkList{Val: data, Next: nil}
	}

	tmp := &LinkList{Val: data, Next: nil}
	tmp.Next = node.Next
	node.Next = tmp

}

func (node *LinkList) tailInsert(data int) {

	if node.Next != nil {
		node.Next.tailInsert(data)
	} else {
		node.Next = &LinkList{Val: data, Next: nil}
	}

}

func headInsertData() {

	head := &LinkList{Val: 0, Next: nil}

	head.headInsert(4)
	head.headInsert(3)
	head.headInsert(2)
	head.headInsert(1)

	fmt.Println("pre reverse")
	LinkPrint(head)

	fmt.Println("after reverse")
	LinkPrint(LinkReverse(head))

}

func tailInsertData() {
	head := &LinkList{Val: 0, Next: nil}

	head.tailInsert(1)
	head.tailInsert(2)
	head.tailInsert(3)
	head.tailInsert(4)

	LinkPrint(head)

}

func LinkPrint(head *LinkList) {

	if head != nil {
		fmt.Printf("head:%v", head.Val)
	}

	for p := head.Next; p != nil; p = p.Next {
		fmt.Printf("->%v", p.Val)
	}

	fmt.Println()
}

func main() {

	headInsertData()

	tailInsertData()

}
