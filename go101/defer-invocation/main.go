package main

type Foo struct {
	v int
}

func MakeFoo(n *int) Foo {
	print("foo")
	print(*n)
	return Foo{}
}

func MakeFoo1(n *int) Foo {
	print("foo1_")
	print(*n)
	return Foo{}
}

func (Foo) Bar(n *int) {
	print("bar")
	print(*n)
}

func main() {
	var x = 1
	var p = &x
	defer MakeFoo1(p).Bar(p) // line 19
	x = 2
	p = new(int) // line 21
	MakeFoo(p)
}
