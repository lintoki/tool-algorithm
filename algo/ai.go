package main

import "fmt"

func aa(a []int) {
	a[0] = 12
	a = append(a, 2)
	a = append(a, 2)
	a = append(a, 2)
}

func main() {
	a := make([]int, 1, 10)
	fmt.Println(a, len(a), cap(a))
	aa(a)
	fmt.Println(a)
}
