package main

import (
	"errors"
	"fmt"
)

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// 函数是第一类型，可作为参数或返回值
func test(x int) func() { // 返回函数类型
	return func() { // 匿名函数
		println(x) // 闭包
	}
}

// defer定义延迟调用，无论函数是否出错，它都确保结束前被调用
func test1(a, b int) {
	defer println("dispose...") // 常用来释放资源，解除锁定，或执行一些清理操作
	println(a / b)
}

func main() {
	a, b := 10, 2       // 定义多个变量
	c, err := div(a, b) // 接收多返回值
	fmt.Println(c, err)
	println("next...")
	x := 100
	f := test(x)
	f()
	println("next...")
	test1(10, 0)
}
