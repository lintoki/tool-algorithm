package main

import (
	"fmt"
)

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

//继承
type dad struct {
	name string
	age  string
}

type son struct {
	dad
	like string
}

func (d dad) say() {
	fmt.Println("i am  dad")
}

//接口
type animal interface {
	Eat()
	Run()
}

type cat struct {
	Age    int
	Weight int
}

type dog struct {
	Age int
}

func (c cat) Eat() {
	fmt.Println("cat eat la")
}

func (c cat) Run() {
	fmt.Println("cat run")
}

func (d dog) Eat() {
	fmt.Println("cat eat la")
}

func (d dog) Run() {
	fmt.Println("cat run")
}

func playPet(a animal) {
	a.Eat()
	a.Run()
}

//自定义类型
type expr interface {
	eval(env Env) float64
}
type va string
type ligle float64
type binary struct {
	op   rune
	x, y expr
}
type Env map[va]float64

func (b binary) eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.eval(env) + b.y.eval(env)
	case '-':
		return b.x.eval(env) - b.y.eval(env)
	default:
		return b.x.eval(env) + b.y.eval(env)
	}
}

func (l ligle) eval(env Env) float64 {
	return float64(l)
}

//类型断言，判断s是否在有s.run方法
func judgeinterface(s interface{}) {
	type a interface {
		run()
	}

	if s, ok := s.(a); ok {
		s.run()
	}

	return
}



func main() {
	//断言
	var x interface{} = 1
	if f, ok := x.(int); ok {
		fmt.Println(f)
	}

	//fib2(5)
	//var i int8 = 127
	//fmt.Println(i,i+1,i+11,i*i)

	//s:= son{dad{"11","22"},"22"}
	//a := new(son)
	//a.say()

	//接口
	//c := cat{1, 1}
	//playPet(c)

}
