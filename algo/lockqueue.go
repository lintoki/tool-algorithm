package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

type lockQueue struct {
	sync.Mutex
	node unsafe.Pointer
}

type lockNode struct {
	val  any
	next unsafe.Pointer
}

func NewLockQueue(val any) *lockQueue {
	head := &lockNode{
		val: val,
	}

	return &lockQueue{
		node: unsafe.Pointer(head),
	}
}

func (q *lockQueue) push(val any) bool {
	defer q.Unlock()
	q.Lock()

	lockNode := &lockNode{
		val: val,
	}

	cur := q.load(&q.node)
	for {
		if cur.next == nil {
			cur.next = unsafe.Pointer(lockNode)
			break
		}

		cur = q.load(&cur.next)
	}

	return true
}

func (q *lockQueue) pop() *lockNode {
	defer q.Unlock()
	q.Lock()

	cur := q.load(&q.node)
	next := q.load(&cur.next)

	for {
		if next.next == nil {
			cur.next = nil
			return next
		}

		cur = q.load(&cur.next)
		next = q.load(&next.next)
	}
}

func (q *lockQueue) load(p *unsafe.Pointer) *lockNode {
	return (*lockNode)(atomic.LoadPointer(p))
}

func (q *lockQueue) cas(p *unsafe.Pointer, old *lockNode, new *lockNode) bool {
	return atomic.CompareAndSwapPointer(p, unsafe.Pointer(old), unsafe.Pointer(new))
}

func (q *lockQueue) syncPush(val any) {
	lockNode := &lockNode{val: val}

	head := q.load(&q.node)
	lockNode.next = unsafe.Pointer(head)
	q.cas(&(q.node), head, lockNode)
}

func (q *lockQueue) syncPop() *lockNode {
	head := q.load(&q.node)
	next := q.load(&head.next)
	q.cas(&(q.node), head, next)

	return head
}

func main() {
	head := NewLockQueue(0)
	for i := 1; i <= 6; i++ {
		a := i
		go func(a int) {
			head.syncPush(a)
		}(a)
	}

	time.Sleep(time.Second * 2)

	fmt.Println(head.syncPop().val)
	fmt.Println(head.syncPop().val)
	fmt.Println(head.syncPop().val)
	fmt.Println(head.syncPop().val)
	fmt.Println(head.syncPop().val)
	fmt.Println(head.syncPop().val)

}
