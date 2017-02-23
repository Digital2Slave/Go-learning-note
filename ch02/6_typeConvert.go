package main

// 除常量、别名类型以及未命名类型外，Go强制要求使用显示类型转换。Go不支持操作符重载！！！

func main() {
	a := 10
	b := byte(a)
	c := a + int(b) // 混合类型表达式必须确保类型一致
	println(c)

	println("next...")

	// 不能将非bool类型结果当作true/false使用
	// x := 100
	// var b bool = x // cannot use x (type int) as type bool in assignment

	println("next...")
	// 如果转换的目标是指针、单向通道或没有返回值的函数类型，那么必须使用括号，以避免造成语法分解错误。
	x := 100
	// p := *int(&x) // cannot convert &x (type *int) to type int
	p := (*int)(&x)
	println(p)
	// (*int)(p)
	// (<-chan int)(c)
	// (func())(x)
	// func() int (x) 仍建议写为 (func()int)(x)
}
