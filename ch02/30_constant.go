package main

import (
	"fmt"
	"unsafe"
)

const (
	ptrSize = unsafe.Sizeof(uintptr(0))
	strSize = len("hello, world!")
)

const x, y int = 123, 0x22
const (
	s    = "hello, world"
	c    = '我'      // rune (unicode code point)
	i, f = 1, 0.123 // int float64 (默认)
	b    = false
)

func main() {
	const x = 123
	println(x)

	const y = 1.23 // 未使用，不会引起编译错误

	{ // 在不同作用域定义同名常量
		const x = "abc"
		println(x)
	}

	const (
		x1, y1 int  = 99, -99
		b      byte = byte(x1) // x被指定为int类型，需显示转换为byte类型
		// n    uint8 = uint8(y) // 错误：constant -999 overflows unit8
	)

	const (
		x2 uint16 = 120
		y2        // 与上一行x类型、右值相同
		s2 = "abc"
		z  // 与s类型、右值相同
	)
	fmt.Printf("%T, %v\n", y2, y2)
	fmt.Printf("%T, %v\n", z, z)

}
