package algorithm

import "fmt"

type LinkNode struct {
	Value interface{}
	Next  *LinkNode
}

type LinkList struct {
	Next *LinkNode
}

func LinkReverse(head *LinkNode) {

}

func (l *LinkList) Insert(data interface{}) {

	if l.Next == nil {
		l.Next = &LinkNode{Value: data, Next: nil}
		return
	}

	l.Next.insert(data)
}

func (n *LinkNode) insert(data interface{}) {

	if n.Next == nil {
		n.Next = &LinkNode{Value: data, Next: nil}
	} else {
		n.Next.insert(data)
	}

}

func (head *LinkNode) ReverseAll() *LinkNode {

	if head.Next == nil {
		return head
	}

	last := head.Next.ReverseAll()
	head.Next.Next = head
	head.Next = nil

	return last
}

var successor *LinkNode

func (head *LinkNode) ReversePreN(m int) *LinkNode {

	if m == 1 {
		successor = head.Next
		return head
	}

	last := head.Next.ReversePreN(m - 1)

	head.Next.Next = head
	head.Next = successor

	return last
}

func (head *LinkNode) ReverseInterval(n, m int) *LinkNode {

	if n == 1 {
		return head.ReversePreN(m)
	}

	head.Next = head.Next.ReverseInterval(n-1, m-1)

	return head
}

/*
* 链表倒数N节点, 快慢指针,
 */
func RemoveBackN(head *LinkNode, n int) *LinkNode {

	dummy := &LinkNode{Next: head} //虚拟节点

	low, fast := dummy, dummy
	//fast先走n+1步，然后一起前进，当fast为nil时，low就来到了倒数第n+1个节点(因为需要是倒数第n+1个节点，才方便删除倒数第n个节点)
	for i := 0; i < n+1; i++ {
		fast = fast.Next //因为n始终合法，所以不需要考虑边界
	}

	for fast != nil {
		fast = fast.Next
		low = low.Next
	}

	//删除倒数第n个节点
	low.Next = low.Next.Next
	return dummy.Next

}

func (head *LinkNode) Print() {

	for p := head; p != nil; p = p.Next {
		fmt.Println(p.Value)
	}

}

func ListReverse(head *LinkNode) *LinkNode {

	var (
		pre  *LinkNode = nil
		cur  *LinkNode = head
		next *LinkNode = head
	)

	/*
	* 需要三个指针, 分别前驱节点, 当前节点后后继节点, 保证交换节点时,链表不断(可以查找到后续节点)
	 */

	for cur != nil {

		next = cur.Next
		// 节点交换, 前一个节点和当前节点交换
		cur.Next = pre
		pre = cur

		cur = next
	}

	return pre
}
