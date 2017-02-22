package main

func main() {
	var x1 int32
	var s = "hello world"
	println(x1, s)
	x := 100
	println(x)
	// if
	if x > 0 {
		println("x")
	} else if x < 0 {
		println("-x")
	} else {
		println("0")
	}
	// switch
	switch {
	case x > 0:
		println("x")
	case x < 0:
		println("-x")
	default:
		println("0")
	}
	// for
	for i := 0; i < 5; i++ {
		println(i)
	}

	for i := 4; i >= 0; i-- {
		println(i)
	}
	println("next...while like")
	// while like
	y := 0
	for y < 5 {
		println(y)
		y++
	}
	println("next...while true like")
	z := 4
	for {
		println(z)
		z--
		if z < 0 {
			break
		}
	}
	println("next...for...range")
	z1 := []int{100, 101, 102}
	for i, n := range z1 {
		println(i, ":", n)
	}

}
