package main

const (
	x = iota // 0
	y        // 1
	z        // 2
)

type color byte

const (
	black color = iota // 指定常量类型
	red
	blue
)

func test(c color) {
	println(c)
}

func main() {
	test(red)
	test(100) // 不能将取值范围限定在预定义的枚举值内
	x1 := 2
	test(color(x1))
	// 数字常量不会分配存储空间，无需像变量那样通过内存寻址来取值，因此无法获取地址.
	// const y = 0x200
	// println(&y) // cannot take the address of y
	const x = 100
	const y byte = x
	// println(x, y)
	const a int = 100
	// const b byte = a // cannot use a (type int) as type byte in const initializer
	// println(a, b)
}
