package main

type link struct {
	data int
	next *link
	len  int
}

func (l *link) new() *link{
	head := new(link)
	head.data = 1
	head.next = nil
	head.len = 1
	return head
}

func (l *link) add(data int) {
	newLink := new(link)
	newLink.next = l

	if l.next == nil {
		l.next = newLink
		l.len++
		return
	}

	cur := l
	for {
		if cur.next == cur {
			cur.next = newLink
			l.len++
		}

		cur = cur.next
	}
}

func main() {
	h := new(link).new()
	
}
