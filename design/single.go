package main

import "sync"

type class struct {
}

var instanse *class
var once sync.Once

func getInstanse() *class {
	if instanse != nil {
		once.Do(func() {
			instanse = &class{}
		})
	}

	return instanse
}

