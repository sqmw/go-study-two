package go_data_structure

import "fmt"

// ListNode 定义一个单链表的 节点
type ListNode[T any] struct {
	Val  T
	next *ListNode[T]
}

// List 定义链表本身
type List[T any] struct {
	head, tail *ListNode[T]
	length     int
}

// TailAdd 定义一个尾部添加的方法
func (listP *List[T]) TailAdd(nodeP *ListNode[T]) {
	if listP.tail == nil {
		listP.tail = nodeP
		listP.head = nodeP
		return
	}
	listP.tail.next = nodeP
	listP.tail = nodeP
}

// 实现 String 方法，但是这里 指针实现的，因此传递的时候必须传递指针
func (listP *List[T]) String() string {
	desStr := "["
	for cur := listP.head; cur != nil; cur = cur.next {
		desStr += fmt.Sprintf("%v, ", cur.Val)
	}
	return desStr + "]"
}
