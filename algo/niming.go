package main

import (
	"fmt"
	"sort"
)

var preregs = map[string][]string{
	"algorithms":            {"data structures"},
	"calculus":              {"linear algebra"},
	"compilers":             {"data structures", "formal languages", "computer organization",},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization",},
	"programming languages": {"data structures", "computer organization"},
}

func fibb(i int) {
	var a = 1
	var x, y = 1, 1
	var dg func(x, y int)
	dg = func(x, y int) {
		fmt.Println(y)
		if a > i {
			return
		}
		a++
		dg(y, x+y)
	}

	dg(x, y)
	return
}

func main() {
	fibb(5)
	//for i, course := range toposort(preregs) {
	//	fmt.Println(i+1, course)
	//}
	//f := F()
	//f[0]()
	//f[1]()
	//f[2]()
}

func F() []func() {
	var b = make([]func(), 5, 5)
	fmt.Println(b)
	for i := 0; i < 5; i++ {
		j := i
		b[i] = func() {
			fmt.Println(j, &j)
		}
	}

	return b
}

func toposort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	fmt.Println(order)

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}
