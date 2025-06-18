package main

import "fmt"

// 声明一种新的数据类型myint，其实就是int的一个别名
type myint int

// 定义一个结构体类型
type book struct {
	title string
	auth string
}

// 传递一个book的副本
func changeBook(b1 book)  {
	b1.auth = "666"
}

// 传递一个book的指针
func changeBook2(b1 *book)  {
	b1.auth = "666"
}

func main() {
	var a myint = 100
	fmt.Println("a = ", a)
	fmt.Printf("a type is %T\n", a)

	var book1 book
	book1.title = "Go语言"
	book1.auth = "www.baidu.com"
	//fmt.Println(book1)

	fmt.Printf("%v\n",book1)

	fmt.Println("---------------------")
	// 打印发现数据未发生任何改变 因为函数传递的是副本
	changeBook(book1)
	fmt.Printf("%v\n",book1)

	fmt.Println("---------------------")
	// 打印发现数据发生改变 因为函数传递的是指针
	// 这里传的是book的地址
	changeBook2(&book1)
	fmt.Printf("%v\n",book1)

}