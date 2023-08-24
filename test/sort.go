package main

import "fmt"

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2

	left := arr[:mid]
	right := arr[mid:]

	return msort(mergeSort(left), mergeSort(right))
}

func msort(left []int, right []int) []int {
	result := make([]int, 0)

	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	if i < len(left) {
		result = append(result, left[i:]...)
	}

	if j < len(right) {
		result = append(result, right[j:]...)
	}

	return result
}

type ListNode struct {
	Val  int       // 节点的值
	Next *ListNode // 指向下一个节点的指针
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head

	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}

	return pre
}

type aa struct {
	v int
	w *aa
}

type monkey struct {
	m    int     // 猴子的编号
	next *monkey // 指向下一个节点的指针
}

// 定义一个函数newMonkey，接受一个整数参数v，返回一个新创建的monkey节点
func newMonkey(v int) *monkey {
	return &monkey{
		m:    v,
		next: nil,
	}
}

// 定义一个函数monkeyKing，接受一个整数参数num，输出最后的大王编号
func monkeyKing(num int) {
	// 创建一个头节点head，并将其指向自己，形成一个只有一个节点的循环链表
	head := newMonkey(1)
	head.next = head
	// 定义一个当前节点cur，并将其初始化为head
	cur := head
	fmt.Println(fmt.Sprintf("%+v,%+v", head, cur))
	// 用一个循环创建剩余的9个节点，并将它们依次插入到循环链表中
	for i := 2; i <= 2; i++ {
		// 创建一个新节点item，并将其编号设为i
		item := &monkey{i, head}
		// 将item插入到cur和cur.next之间，并更新cur为item
		fmt.Println(fmt.Sprintf("%+v,%+v", head, cur))
		cur.next = item
		fmt.Println(fmt.Sprintf("%+v,%+v", head, cur))
		cur = item
		fmt.Println(fmt.Sprintf("%+v,%+v", head, cur))
	}
}

func main() {
	//var arr = []int{9, 4, 5, 7, 8, 8, 8, 9, 9, 97, 5, 3, 2342, 4, 23, 34, 5, 4, 45, 6, 567, 65, 756, 8, 76, 867, 78, 9, 78}
	//fmt.Println(mergeSort(arr))

	a := &aa{1, nil}
	a.w = a
	b := a
	fmt.Println(fmt.Sprintf("%+v,%+v", a, b))
	c := &aa{3, nil}
	b.v = 2
	b.w = c
	fmt.Println(fmt.Sprintf("%+v,%+v", a, b))
	b = c
	fmt.Println(fmt.Sprintf("%+v,%+v", a, b))

	monkeyKing(2)
}
