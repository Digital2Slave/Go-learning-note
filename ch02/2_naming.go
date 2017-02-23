package main

import (
	"strconv"
)

/*
* 1. 以字母或下划线开始，由多个字母、数字和下划线组合而成
* 2. 区分大小写
* 3. 使用驼峰(camel case)拼写格式
* 4. 局部变量优先使用短命
* 5. 不用使用保留关键字
* 6. 不建议使用与预定义常量、类型、内置函数相同的名字
* 7. 专有名词通常会全部大写，例如escapeHTML
* 8. 符号名字首字母大写的为导出成员，可被包外引用；小写的仅能在包内使用。
* 空标识符"_"，通常作为忽略占位符使用，可作表达式左值，无法读取内容，是预置成员！
 */

func main() {
	var c int                 // c代替count
	for i := 0; i < 10; i++ { // i代替index
		c++
	}
	println(c)

	x, _ := strconv.Atoi("12") // "_"忽略Atoi的err返回值
	println(x)
}
