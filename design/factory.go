package main

import "fmt"

type tool interface {
	use(s string)
}
type factory struct{}

type pen struct{}

func (p1 *pen) use(s string) {
	fmt.Println(s)
}

type pencial struct{}

func (p1 *pencial) use(s string) {
	fmt.Println(s)
}

func (f *factory) New(s string) tool {
	switch s {
	case "pen":
		return &pen{}
	case "pencial":
		return &pencial{}
	default:
		return nil
	}
}

func main() {
	f := factory{}
	f.New("pen").use("sadasd")
}
