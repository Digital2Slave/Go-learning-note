package main

import (
	"fmt"
)

type X int

func (x *X) inc() { // 名称钱的参数称作receiver,作用类似python self
	*x++
}

// 直接调用匿名字段的方法，这种方式可实现和继承类似的功能
type user struct {
	name string
	age  byte
}

func (u user) ToString() string {
	return fmt.Sprintf("%+v", u)
}

type manager struct {
	user
	title string
}

func main() {
	var x X
	x.inc()
	println(x)
	println("next...")
	var m manager
	m.name = "Tom"
	m.age = 29
	println(m.ToString()) // 调用user.ToString()
}
