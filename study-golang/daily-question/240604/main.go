package main

func main() {
	v := new(int)
	*v = 2
	// *v 2
	// - 表示求负数，+ 表示 0 +
	// 5 / 0 + (-2)
	// 5 / -2 = -2
	println(5 / +-*v)
}
