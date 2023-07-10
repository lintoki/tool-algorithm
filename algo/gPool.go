package main

import (
	"fmt"
	"sync"
)

// Pool 是一个表示协程池的结构体
type Pool struct {
	size    int            // 协程池的大小
	tasks   chan func()    // 接收任务的通道
	wg      sync.WaitGroup // 同步协程的等待组
	running bool           // 表示协程池是否运行的标志
}

// NewPool 创建一个新的协程池，指定大小
func NewPool(size int) *Pool {
	return &Pool{
		size:    size,
		tasks:   make(chan func()),
		running: false,
	}
}

// Start 启动协程池，运行协程
func (p *Pool) Start() {
	if p.running {
		return
	}
	p.running = true
	for i := 0; i < p.size; i++ {
		p.wg.Add(1)
		go p.worker()
	}
}

// Stop 停止协程池，等待协程结束
func (p *Pool) Stop() {
	if !p.running {
		return
	}
	p.running = false
	close(p.tasks)
	p.wg.Wait()
}

// Submit 提交一个任务到协程池
func (p *Pool) Submit(task func()) {
	if !p.running {
		return
	}
	p.tasks <- task
}

// worker 是每个协程运行的函数
func (p *Pool) worker() {
	defer p.wg.Done()
	for task := range p.tasks {
		task()
	}
}

// example task that prints a message
// 示例任务，打印一条消息
func task(msg string) func() {
	return func() {
		fmt.Println(msg)
	}
}

func main() {
	// 创建一个有10个协程的协程池
	pool := NewPool(10)

	// 启动协程池
	pool.Start()

	// 提交一些任务到协程池
	for i := 0; i < 100; i++ {
		pool.Submit(task(fmt.Sprintf("task %d", i)))
	}

	// 停止协程池，等待任务完成
	pool.Stop()
}
