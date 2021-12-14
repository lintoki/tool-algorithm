package main

import "fmt"

//斐波那契额数列
func fib(n int) {
	var x, y = 0, 1
	fmt.Println(x)

	for i := 0; i <= n; i++ {
		x, y = y, x+y
		fmt.Println(x)
	}
	return
}

//斐波那契数列递归
func fib2(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(fibcourson(i))
	}
	return
}

func fibcourson(n int) int {
	if n < 2 {
		return n
	}
	return fibcourson(n-1) + fibcourson(n-2)
}

func main() {
	//fib2(5)
	var i int8 = 127
	fmt.Println(i,i+1,i+11,i*i)
}
