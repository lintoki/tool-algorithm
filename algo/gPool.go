package main

import (
	"fmt"
	"sync"
)

type gPool struct {
	size      int
	worker    chan func()
	close     []chan int
	runWorker int
	wg        *sync.WaitGroup
}

func NewGPool(size int) *gPool {
	g := &gPool{
		size:      size,
		worker:    make(chan func(), size*2),
		close:     make([]chan int, 0),
		runWorker: 0,
		wg:        &sync.WaitGroup{},
	}

	return g
}

func (g gPool) run() {
	for i := 0; i < g.size; i++ {
		g.wg.Add(1)
		go g.work()
	}

	close(g.worker)
	g.wg.Wait()
}

func (g gPool) submit(fn func()) string {
	select {
	case g.worker <- fn:
		return "success"
	default:
		return "pool is full"
	}

}

func (g gPool) work() {
	defer g.wg.Done()

	for {
		if fn, ok := <-g.worker; ok {
			fn()
		} else {
			return
		}
	}
}

func task(msg int) func() {
	return func() {
		fmt.Println(fmt.Sprintf("task 1 :%v", msg))
	}
}

func task2(msg int) func() {
	return func() {
		fmt.Println(fmt.Sprintf("task 2 :%v", msg))
	}
}

func main() {
	p := NewGPool(2)
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(p.submit(task(i)))
		} else {
			fmt.Println(p.submit(task2(i)))
		}

	}

	p.run()
}
