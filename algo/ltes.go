package main

import "fmt"

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
	i:=1
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

func main() {
	head := new(tl).newt(1)
	for i := 2; i < 5; i++ {
		head.addSingle(i)
	}

	head.list()

	//for i := 2; i < 5; i++ {
	//	head.add(i)
	//}
	//
	//head.list()
	//
	//getKing(head, 3)

	//h2 := new(tl).newt(1)
}
