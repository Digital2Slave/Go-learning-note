package main

import (
	"fmt"
	"math"
	"strconv" // 不同进制(字符串)间转换
)

func test(x byte) {
	println(x)
}

func main() {
	a, b, c := 100, 0144, 0x64
	fmt.Println(a, b, c)
	fmt.Printf("0b%b, %#o, %#x\n", a, b, c)
	fmt.Println(math.MinInt8, math.MaxInt8)

	println("next...")

	a1, _ := strconv.ParseInt("1100100", 2, 32)
	b1, _ := strconv.ParseInt("0144", 8, 32)
	c1, _ := strconv.ParseInt("64", 16, 32)
	println(a1, b1, c1)
	println("0b" + strconv.FormatInt(a1, 2))
	println("O" + strconv.FormatInt(b1, 8))
	println("Ox" + strconv.FormatInt(c1, 16))

	println("next...")
	var a2 float32 = 1.1234567899 // 注意：默认浮点类型是float64
	var b2 float32 = 1.12345678
	var c2 float32 = 1.123456781
	println(a2, b2, c2)
	println(a == b, a == c)
	fmt.Printf("%v %v, %v", a2, b2, c2)

	println("next...alias")
	// byte alias for uint8
	// rune alias for int32
	var a3 byte = 0x11
	var b3 uint8 = a3
	var c3 uint8 = a3 + b3
	test(c3)
}
