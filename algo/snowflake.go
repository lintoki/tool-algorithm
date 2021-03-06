package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

const (
	nodeBits  uint8 = 10                    // 节点 ID 的位数
	stepBits  uint8 = 12                    // 序列号的位数
	nodeMax   int64 = -1 ^ (-1 << nodeBits) // 节点 ID 的最大值，用于检测溢出
	stepMax   int64 = -1 ^ (-1 << stepBits) // 序列号的最大值，用于检测溢出
	timeShift uint8 = nodeBits + stepBits   // 时间戳向左的偏移量
	nodeShift uint8 = stepBits              // 节点 ID 向左的偏移量
)

var Epoch int64 = 1288834974657 // timestamp 2006-03-21:20:50:14 GMT
type ID int64

// A Snowflake struct holds the basic information needed for a snowflake generator worker
type Node struct {
	mu        sync.Mutex // 添加互斥锁，保证并发安全
	timestamp int64      // 时间戳部分
	node      int64      // 节点 ID 部分
	step      int64      // 序列号 ID 部分
}

// NewNode returns a new snowflake worker that can be used to generate snowflake IDs
func NewNode(node int64) (*Node, error) {
	// 如果超出节点的最大范围，产生一个 error
	if node < 0 || node > nodeMax {
		return nil, errors.New("Node number must be between 0 and 1023")
	}
	// 生成并返回节点实例的指针
	return &Node{
		timestamp: 0,
		node:      node,
		step:      0,
	}, nil
}

// Generate creates and returns a unique snowflake ID
func (n *Node) Generate() ID {

	n.mu.Lock()         // 保证并发安全, 加锁
	defer n.mu.Unlock() // 方法运行完毕后解锁

	// 获取当前时间的时间戳 (毫秒数显示)
	now := time.Now().UnixNano() / 1e6

	if n.timestamp == now {
		// step 步进 1
		n.step++

		// 当前 step 用完
		if n.step > stepMax {
			// 等待本毫秒结束
			for now <= n.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}

	} else {
		// 本毫秒内 step 用完
		n.step = 0
	}

	n.timestamp = now
	// 移位运算，生产最终 ID
	result := ID((now-Epoch)<<timeShift | (n.node << nodeShift) | (n.step))

	return result
}

func main() {
	// 测试脚本
	fmt.Println((-1 ^ (-1 << 10)))

	// 生成节点实例
	node, err := NewNode(1)

	if err != nil {
		fmt.Println(err)
		return
	}

	ch := make(chan ID)
	count := 10000
	// 并发 count 个 goroutine 进行 snowflake ID 生成
	for i := 0; i < count; i++ {
		go func() {
			id := node.Generate()
			ch <- id
		}()
	}

	defer close(ch)

	m := make(map[ID]int)
	for i := 0; i < count; i++ {
		id := <-ch
		// 如果 map 中存在为 id 的 key, 说明生成的 snowflake ID 有重复
		_, ok := m[id]
		if ok {
			fmt.Printf("ID is not unique!\n")
			return
		}
		// 将 id 作为 key 存入 map
		m[id] = i
	}
	// 成功生成 snowflake ID
	fmt.Println("All ", count, " snowflake ID generate successed!\n")
}
