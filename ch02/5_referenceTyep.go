package main

// 所谓引用类型(reference type) 特指slice、map、channel这三种预定义类型。
// 必须使用make函数创建，编译器会将make转换为目标类型专用的创建函数(或指令)，以确保完成全内存分配和相关属性初始化。

func mkslice() []int {
	s := make([]int, 0, 10)
	s = append(s, 100)
	return s
}

func mkmap() map[string]int {
	m := make(map[string]int)
	m["a"] = 1
	return m
}

func main() {
	m := mkmap()
	println(m["a"])

	s := mkslice()
	println(s[0])
}
