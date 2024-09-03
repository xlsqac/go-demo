// package main
// 方法
package main

import "fmt"

type Book struct {
	Pages int
}

// 值类型的属主参数
func (b Book) setPages(page int) {
	b.Pages = page
}

// 指针类型的属主参数
func (b *Book) setPages1(page int) {
	b.Pages = page
}

func main() {
	// 简单测试一下不同类型的属主参数，能否互相调用，以及调用能否正常修改字段
	{
		b := Book{10}
		b1 := &Book{20}

		b.setPages(100)
		fmt.Println(b.Pages) // 不可以修改

		b.setPages1(1000)
		fmt.Println(b.Pages) //  可以修改

		b1.setPages(100)
		fmt.Println(b1.Pages) // 不可以修改
	}

}
