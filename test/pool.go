package main

import (
	"fmt"
	"sync"
	"time"
)

type pool struct {
	size   int
	worker chan func()
	wg     *sync.WaitGroup
}

func newPool(size int) *pool {
	return &pool{
		size:   size,
		worker: make(chan func(), size),
		wg:     &sync.WaitGroup{},
	}
}

func (p *pool) run() { // 使用指针接收者
	for i := 0; i < p.size; i++ {
		p.wg.Add(1)
		go p.work()
	}
}

func (p *pool) work() { // 使用指针接收者
	defer p.wg.Done()
	n := 0
	for {
		select {
		case fn, _ := <-p.worker:
			fn()
			n = 0
			continue
		default:
			if n == 10 {
				fmt.Println(fmt.Sprintf("nis ========== %v", n))
				return
			} else {
				n++
				time.Sleep(time.Second * 1)
				continue
			}

		}
	}
}

func (p *pool) submit(fn func()) bool { // 使用指针接收者
	select {
	case p.worker <- fn:
		return true
	default:
		return false
	}

}

func t1() {
	fmt.Println("func t111")
	return
}

func t2() {
	fmt.Println("func t222")
	return
}

func exm() {
	flag := true
	p := newPool(5)
	p.run()
	defer close(p.worker) // 使用 defer 关闭 worker 通道
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			flag = p.submit(t1)
		} else {
			flag = p.submit(t2)
		}
		if flag == true {
			fmt.Println(fmt.Sprintf("func %v,is success", i))
		} else {
			fmt.Println(fmt.Sprintf("func %v,is full", i))
		}
	}

	p.wg.Wait()
}

func main() {
	exm()
}
