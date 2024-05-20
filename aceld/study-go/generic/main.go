package main

func plus[T int | uint](n1, n2 T) T {
	return n1 + n2
}

func main() {
	n1, n2 := 1, 2
	plus(n1, n2)

	u1, u2 := uint(1), uint(2)
	plus(u1, u2)

}
