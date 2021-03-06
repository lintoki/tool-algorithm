package main

import (
	"fmt"
	"time"
)

//无缓冲通道
func c1() {
	done := make(chan struct{})
	fmt.Println(1)

	go func() {
		fmt.Println(2)
		done <- struct{}{}
	}()

	fmt.Println(3, <-done)

	fmt.Println(4)
}

//管道
func couter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func squer(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * x
	}

	close(out)
}

func printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}

}

func c2() {
	num := make(chan int)
	sqn := make(chan int)

	go couter(num)
	go squer(num, sqn)
	printer(sqn)
}

//select
func sel() {
	nozs := make(chan int64, 1)

	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println(time.Now().Unix())
		nozs <- time.Now().Unix()
	}()

	select {
	case x := <-nozs:
		fmt.Println(x)
	case <-time.After(time.Second * 7):
		fmt.Println(22)
		//default:
		//	fmt.Println(777)
	}

	return
}

func main() {
	//无缓冲通道
	//c1()

	//管道
	//c2()

	//select
	sel()
}
