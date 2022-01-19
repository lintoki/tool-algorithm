package main

import "fmt"

type link struct {
	data int
	next *link
	len  int
}

func (l *link) new(data int) *link {
	head := new(link)
	head.data = data
	head.next = nil
	head.len = 1
	return head
}

func (l *link) add(data int) {
	newLink := l.new(data)
	newLink.next = l

	if l.next == nil {
		l.next = newLink
		l.len++
		return
	}

	cur := l
	for {
		if cur.next == l {
			cur.next = newLink
			l.len++
			return
		}

		cur = cur.next
	}
}

func (l *link) list() {
	tmp := l
	for {
		fmt.Println(tmp.data)
		if tmp.next == l {
			return
		}

		tmp = tmp.next
	}
}

func main() {
	h := new(link).new(1)
	for i := 2; i < 6; i++ {
		h.add(i)
	}

	h.list()

	cur := h
	i := 1
	for {
		if cur.next == cur {
			fmt.Println(cur.data)
			return
		}

		if i == 2 {
			i = 1
			cur.next = cur.next.next
		} else {
			i++
		}

		cur = cur.next
	}
}
