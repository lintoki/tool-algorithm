package main

import (
	"fmt"
	"reflect"
)

func main(){
	var a interface{}
	a =122

	// 从接口值到反射对象
	fmt.Println(reflect.TypeOf(a),reflect.ValueOf(a))

	//从反射对象到接口值
	var f float64 = 3.1415
	v := reflect.ValueOf(f)   // f 隐式地被转成了interface{}
	y := v.Interface().(float64)
	fmt.Println(v,y)

	//想要修改一个反射对象，那么该值必须是可以被设置的，必须传递指针，而不是值
	var f float64 = 3.1415
	v := reflect.ValueOf(&f).Elem()   // 传了f的指针，&f 隐式地被转成了interface{}
	v.SetFloat(2.873)
	fmt.Println(f,v)

}