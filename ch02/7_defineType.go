package main

import (
	"fmt"
)

// 使用关键字type定义用户自定义类型，包括基于现有基础类型创建，或者结构体、函数类型等。

type flags byte

const (
	read  flags = 1 << iota // 1
	write                   // 2
	exec                    // 4
)

func main() {
	println(read, write, exec)
	f0 := read | exec // 2
	fmt.Printf("%b\n", f0)

	println("next...")
	// 和var、const类似，多个type定义可合并成组，可在函数或代码内定义局部类型。
	// 自定义类型即便指定了基础类型，除操作符外，它不会继承基础类型的其他信息(包括方法)。不能视作别名，不能隐式转换，不能直接用于表达式。
	type ( // 组
		user struct { // 结构体
			name string
			age  uint8
		}
		event func(string) bool // 函数类型
	)

	u := user{"Tome", 20}
	fmt.Println(u)

	var f event = func(s string) bool {
		println(s)
		return s != ""
	}
	f("abc")
}
