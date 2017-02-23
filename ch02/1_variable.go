package main

import (
	"fmt"
)

var (
	x, y int
	a, s = 100, "abc"
)

func main() {
	fmt.Println("Hello, go!")
	println(&x, ":", x, y, &a, ":", a, &s, ":", s)

	println("next...")
	x := 100
	a, s := 1, "xyz"
	println(&x, ":", x, &a, ":", a, &s, ":", s)

	println("next...")
	m := 100
	println(&m, ":", m)
	m, n := 200, 300 // m 退化为赋值操作。最少有一个新变量被定义，且必须是同一个作用域！
	println(&m, ":", m, n)

	println("next...")
	p, q := 1, 2
	p, q = q+3, p+2
	println(p, q)
	// 编译器将未使用局部变量当做错误。
}
