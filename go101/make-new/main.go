// package main
// make 和 new 函数的区别
package main

import "fmt"

func myMake() {
	// make 一个初始长度为 10 的slice
	s := make([]int, 10)
	fmt.Println(s)

	// make 一个容量为 10 的 map
	m := make(map[string]int, 10)
	m["k1"] = 1
	fmt.Println(m)

	// make 一个无缓冲的 channel
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	fmt.Println(<-ch)
}

func myNew() {
	// new 一个 int 类型的指针
	i := new(int)
	fmt.Println("pointer addr:", i, "value:", *i)

	type S struct {
		Name string
	}
	// new 一个 S 类型的指针
	s := new(S)
	s.Name = "S"
	fmt.Println("pointer addr:", s, "value:", *s)

	var a *int
	a = new(int)
	*a = 1

	//var mv *map[string]string
	//fmt.Printf("mv: %p %#v \n", &mv, mv) //mv: 0xc042004028 (*map[string]string)(nil)
	mv := new(map[string]string)
	fmt.Printf("mv: %p %#v \n", &mv, mv) //mv: 0xc000006028 &map[string]string(nil)
	*mv = make(map[string]string)
	fmt.Printf("mv: %p %#v \n", &mv, mv) //mv: 0xc042004028 &map[string]string{"a":"a"}
	(*mv)["a"] = "a"
	fmt.Printf("mv: %p %#v \n", &mv, mv) //mv: 0xc042004028 &map[string]string{"a":"a"}

	cv := new(chan string)
	fmt.Println(*cv)

	array := new([5]int)
	fmt.Println(*array)

	slice := new([]int)
	if *slice == nil {
		fmt.Println("nil slice")
	}

	m := new(map[string]int)
	if *m == nil {
		fmt.Println("nil map")
	}
}

func main() {
	myMake()
	myNew()

}
