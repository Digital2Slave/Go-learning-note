package main

import (
	"fmt"
)

type user struct { //结构体类型
	name string
	age  byte
}

type manager struct {
	user  // 匿名嵌入其他类型
	title string
}

func main() {
	x := make([]int, 0, 5) // 创建容量为5的切片
	for i := 0; i < 8; i++ {
		x = append(x, i) // 追加数据，当超过容量限制时，自动分配更大的存储空间
	}
	fmt.Println(x)
	println("next...")
	m := make(map[string]int) // 创建字典类型对象
	m["a"] = 1
	y, ok := m["b"] // 使用ok-idiom获取，可知道key/value是否存在
	fmt.Println(y, ok)
	delete(m, "a") // 删除
	println("next...struct")
	var m1 manager
	m1.name = "Tom" // 直接访问匿名字段的成员
	m1.age = 29
	m1.title = "CTO"
	fmt.Println(m1)
}
