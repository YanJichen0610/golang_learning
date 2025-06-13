/*
	四种变量的声明方式
*/
package main

import "fmt"

// 声明全局变量方法一二三是可以的， 但是方法四是不可以的


func main() {
	// 方法一: 声明一个变量 默认值为0
	var a int
	fmt.Println("a = ", a)

	// 方法二: 声明一个变量, 初始化一个默认值
	var b int = 100
	fmt.Println("b = ", b)

	// 方法三: 初始化的时候可以省去数据类型，通过数值自动判断数据的数据类型
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	// 方法四: (常用方法) 省区var关键词, 直接使用自动匹配
	e := 100
	fmt.Println("e = ", e)

	// 声明多个变量方法一 (单行)
	var xx, yy int = 100, 200
	fmt.Println("xx = ", xx,"yy = ", yy)

	// 声明多个变量方法二 (多行)
	var (
		vv int = 100
		jj bool = true
	)
	fmt.Println("vv = ", vv,"jj = ", jj)
}