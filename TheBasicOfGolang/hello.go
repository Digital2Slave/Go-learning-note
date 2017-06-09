package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hello, world or 你好，世界 or καλημ ́ρα κóσμ or こんにちはせかい\n")
	var c complex64 = 5 + 5i
	fmt.Printf("Value is : %v\n", c)
	// change value in string
	s := "hello"
	c1 := []byte(s) // 将字符串 s 转换为 []byte 类型
	c1[0] = 'c'
	s2 := string(c1) // 再转换回 string 类型
	fmt.Printf("%s\n", s2)

	s = "c" + s[1:] // 字符串虽不能更改，但可进行切片操作
	fmt.Printf("%s\n", s)
}
