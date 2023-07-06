package main

import (
	"fmt"
	"testing"
)

type tl struct {
	data int
	next *tl
}

func (t *tl) newt(data int) *tl {
	head := new(tl)
	head.data = data
	head.next = nil
	return head
}

func (t *tl) add(data int) {
	node := t.newt(data)
	node.next = t

	if t.next == nil {
		t.next = node
		return
	}

	cur := t
	for {
		if cur.next == t {
			cur.next = node
			return
		}

		cur = cur.next
	}
}

func getKing(l *tl, num int) {
	cur := l

	i := 1
	for {
		if cur.next == cur {
			fmt.Println(cur.data)
			return
		}

		if i == num-1 {
			cur.next = cur.next.next
			i = 1
		} else {
			i++
		}

		cur = cur.next

	}
}

func (l *tl) list() {
	tmp := l
	for {
		fmt.Println(tmp.data)
		if tmp.next == l || tmp.next == nil {
			return
		}

		tmp = tmp.next
	}
}

func revList(l *tl) {
	cur := l

	var tmp *tl
	//var prev *tl //记录上一次处理的节点，最后一次要用
	i := 1
	for {

		if cur.next == nil {
			cur.next = tmp
			return
		}

		tmp = cur
		cur, cur.next = cur.next, cur
		cur = tmp.next

		i++
	}
}

func (l *tl) addSingle(data int) {
	node := new(tl).newt(data)
	cur := l
	for {
		if cur.next == nil {
			cur.next = node
			return
		}

		cur = cur.next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var rev *ListNode
	rev = nil
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = rev
		rev = cur
		cur = next
	}

	return rev
}

func TestReverseList(t *testing.T) {
	// 构造一个链表 1 -> 2 -> 3 -> 4 -> 5
	head := &ListNode{Val: 1}
	node1 := &ListNode{Val: 2}
	node2 := &ListNode{Val: 3}
	node3 := &ListNode{Val: 4}
	node4 := &ListNode{Val: 5}
	head.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	// 打印原始链表所有节点的值
	fmt.Print("original list: ")
	for p := head; p != nil; p = p.Next {
		fmt.Printf("%d ", p.Val)
	}
	fmt.Println()

	// 反转链表
	reversedHead := reverseList(head)

	// 打印反转后的链表所有节点的值
	fmt.Print("reversed list: ")
	for p := reversedHead; p != nil; p = p.Next {
		fmt.Printf("%d ", p.Val)
	}
	fmt.Println()

	// 验证反转结果是否正确
	if reversedHead.Val != 5 {
		t.Errorf("reversed head value is %d, want %d", reversedHead.Val, 5)
	}
	if reversedHead.Next.Val != 4 {
		t.Errorf("reversed head next value is %d, want %d", reversedHead.Next.Val, 4)
	}
	if reversedHead.Next.Next.Next.Val != 2 {
		t.Errorf("reversed head next next next value is %d, want %d", reversedHead.Next.Next.Next.Val, 2)
	}
}

func main() {
	a := testing.T{}
	TestReverseList(&a)
	//head := new(tl).newt(1)
	//for i := 2; i < 5; i++ {
	//	head.addSingle(i)
	//}
	//
	//head.list()

	//for i := 2; i < 5; i++ {
	//	head.add(i)
	//}
	//
	//head.list()
	//
	//getKing(head, 3)

	//h2 := new(tl).newt(1)
}
