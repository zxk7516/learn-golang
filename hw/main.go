package main

import (
	"fmt"
	"zxk/hw/lru"
)

func main() {
	res := lru.LRU(
		[][]int{{1, 1, 1}, {1, 2, 2}, {1, 3, 2}, {2, 1}, {1, 4, 4}, {2, 2}},
		3,
	)
	fmt.Println(res)

	// m := 1
	// n := 4
	// root := new(ListNode) //头结点
	// p := root
	// for i := 1; i < 7; i++ { //初始化链表
	// 	node := new(ListNode)
	// 	node.Val = i
	// 	p.Next = node
	// 	p = p.Next
	// }
	// p.Next = new(ListNode) //尾节点
	// l1 := root
	// var l2, l3 *ListNode
	// p = root
	// index := 0
	// for l3 == nil { //截成三段
	// 	if index == m {
	// 		l2 = p.Next
	// 		p.Next = nil
	// 		p = l2
	// 	}
	// 	if index == n {
	// 		l3 = p.Next
	// 		p.Next = nil
	// 	}
	// 	p = p.Next
	// 	index++
	// }
	// l2 = reverseList(l2) // 反转l2
	// list := []*ListNode{l2, l3}
	// p = l1
	// index = 0
	// for index < len(list) { //拼接起来
	// 	for p.Next != nil {
	// 		p = p.Next
	// 	}
	// 	p.Next = list[index]
	// 	index++
	// }
	// l1.Print()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode = nil
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre
		// var next *ListNode = cur.Next
		// cur.Next = pre
		// pre = cur
		// cur = next
	}
	return pre
}

func (l *ListNode) Print() {
	for l != nil {
		if l.Val > 0 {
			fmt.Println(l.Val)
		}
		l = l.Next
	}

}
