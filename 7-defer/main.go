package main

import "fmt"

func func1() {
	fnt.Println("func1")
}

func func2() {
	fnt.Println("func2")
}

func func3() {
	fnt.Println("func3")
}

func main() {
	// 知识点1: 写入defer关键字
	// defer fmt.Println("main end1")
	// defer fmt.Println("main end2

	// fmt.Println("main:hello go 1")
	// fmt.Println("main:hello go 2")

	// 知识点2: 执行顺序：
	// 1. 先执行main函数，然后再执行defer语句
	// 2. defer语句的执行顺序是后进先出
	// 3. 所以应该执行defer fmt.Println("func3") -> defer fmt.Println("func2") -> defer fmt.Println("func1")
	// defer func1()
	// defer func2()
	// defer func3()

	// 知识点3: defer和return的执行顺序
	// 先执行return语句，然后再执行defer语句
}