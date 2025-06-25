package main

import (
	"fmt"
)

// interface{} 是万能数据类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called")
	fmt.Println(arg)

	// interface 如何区分此时传入的底层数据类型是什么
	// 给interface{} 提供断言机制
	value, ok := arg.(string)
	// 如果当前数据类型为string 那么arg则为True
	if !ok {
		fmt.Println("arg is not int")
	} else {
		fmt.Println("arg is int", value)
	}
}

type Book struct {
	auth string
}


func main() {
	book := Book{"Golang"}

	// 该方法传入参数可以是任何类型
	myFunc(book)
	fmt.Println("-----------------------")
	myFunc(100)
	fmt.Println("-----------------------")
	myFunc("hello")
}