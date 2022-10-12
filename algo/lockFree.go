package main

import (
	"fmt"
	"sync/atomic"
	"time"
	"unsafe"
)

// LockFreeQueue 使用atomic使用无锁队列
type LockFreeQueue struct {
	head unsafe.Pointer
}

type node struct {
	next  unsafe.Pointer
	value interface{}
}

func NewLockFreeQueue(v interface{}) *LockFreeQueue {
	node := &node{value: v}
	return &LockFreeQueue{
		head: unsafe.Pointer(node),
	}
}

func (l *LockFreeQueue) push(v interface{}) bool {
	for {
		head := atomic.LoadPointer(&l.head)
		node := &node{value: v, next: head}
		new := unsafe.Pointer(node)

		return atomic.CompareAndSwapPointer(&l.head, head, new)
	}
}

func (l *LockFreeQueue) pop() interface{} {
	head := atomic.LoadPointer(&l.head)
	if head == nil {
		return nil
	}

	next := atomic.LoadPointer(&(*node)(head).next)
	if atomic.CompareAndSwapPointer(&l.head, head, next) {
		return (*node)(head).value
	}

	return nil
}

func main() {
	lfQueue := NewLockFreeQueue(1)

	for i := 2; i <= 10; i++ {
		go lfQueue.push(i)
	}

	time.Sleep(time.Second * 1)

	for i := 1; i <= 11; i++ {
		fmt.Println(lfQueue.pop())
	}
}
