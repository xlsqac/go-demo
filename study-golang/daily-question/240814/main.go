package main

import "fmt"

const (
	z = "zz"
	k
)

func main() {
	{
		const (
			a = iota
		)
		fmt.Println(z, k)
	}

	const (
		a = iota // ok
		b
	)
	fmt.Println(a, b)
	const (
		c = iota
	)
	fmt.Println(c)

	{
		const (
			a = "zz"
			b
			c
		)
		fmt.Println(a, b, c)
	}
	{
		const (
			a, b = "golang", 100
			d, e
			f bool = true
			g
		)
		fmt.Println(d, e, g)
	}
	{
		const ()
	}
}
