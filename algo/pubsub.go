package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type (
	subscriber chan interface{}        //订阅者为一个通道
	topicfunc func(v interface{}) bool //主题为一个过滤器
)

type Publisher struct {
	m          sync.RWMutex             //读写锁
	buffer     int                      //订阅队列的缓存大小
	timeout    time.Duration            //发布超时时间
	subscibers map[subscriber]topicfunc //订阅者信息
}

//构建一个发布对象，设置超时时间和缓存队列长度
func NewPublisher(publishTimeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer:     buffer,
		timeout:    publishTimeout,
		subscibers: map[subscriber]topicfunc{},
	}
}

//添加一个新的订阅者，订阅全部主题
func (p *Publisher) Subscribe() chan interface{} {
	return p.SubscribeTopic(nil)

}

//添加一个新的订阅者，订阅过滤器筛选后的主题
func (p *Publisher) SubscribeTopic(topic topicfunc) chan interface{} {
	ch := make(chan interface{}, p.buffer)
	p.m.Lock()
	p.subscibers[ch] = topic
	p.m.Unlock()
	return ch
}

//退出订阅
func (p *Publisher) Ezvict(sub chan interface{}) {
	p.m.Lock()
	defer p.m.Unlock()

	delete(p.subscibers, sub)
	close(sub)
}

//发布一个主题
func (p *Publisher) Publish(v interface{}) {
	p.m.Lock()
	defer p.m.Unlock()

	var wg sync.WaitGroup
	for sub, topic := range p.subscibers {
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)

	}
	wg.Wait()
}

//发布主题，忍耐一定超时
func (p *Publisher) sendTopic(sub subscriber, topic topicfunc, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	if topic != nil && !topic(v) {
		return
	}

	select {
	case sub <- v:
	case <-time.After(p.timeout):
	}
}

//关闭主题
func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	for sub := range p.subscibers {
		delete(p.subscibers, sub)
		close(sub)
	}
}

//ceshi
func main() {
	p := NewPublisher(100*time.Millisecond, 10)
	defer p.Close()

	all := p.Subscribe()

	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}

		return false
	})

	p.Publish("hallo world")
	p.Publish("hallo golang")

	go func() {
		for msg := range all {
			fmt.Println("all:", msg)
		}
	}()

	go func() {
		for msg := range golang {
			fmt.Println("golang:", msg)
		}
	}()

	time.Sleep(3 * time.Second)
}
