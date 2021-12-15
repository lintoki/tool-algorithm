package main

import (
	"fmt"
	"sync"
	"time"
)

type queue struct {
	list []int
	sync.RWMutex
}

func newQueue() *queue {
	return &queue{
		list: []int{},
	}
}

func (q *queue) push(k int) { //从map中读取一个值
	q.Lock()
	defer q.Unlock()
	q.list = append(q.list, k)
}

func (q *queue) pop() { //从map中读取一个值
	q.Lock()
	defer q.Unlock()
	q.list = q.list[:]
}

func main() {
	queue := newQueue()
	for i := 1; i < 10; i++ {
		go queue.push(i)
	}
	time.Sleep(time.Second*2)
	fmt.Println(queue.list)
	queue.pop()
	fmt.Println(queue.list)
	queue.push(55)
	fmt.Println(queue.list)

	//a := []int{1, 2, 3} // ... 会自动计算数组长度
	//b := a
	//a[0] = 100

	//fmt.Println(a,b)

	//intarr := []int{12, 34, 55, 66, 43}
	//slice := intarr[0:4]
	//
	//fmt.Printf("address of slice %p add of Arr %p \n", &slice, &intarr)
	//fmt.Printf("address of slice %p add of Arr %p \n", slice, intarr)
	//intarr = []int{1, 2}
	//fmt.Printf("address of slice %p add of Arr %p \n", &slice, &intarr)
	//fmt.Printf("address of slice %p add of Arr %p \n", slice, intarr)

}

func printLenCap(nums []int) {
	fmt.Printf("len: %d, cap: %d %v\n", len(nums), cap(nums), nums)
}

func test(i []int) {
	//i = []int{1, 2, 3}
	i[2] = 2
}
