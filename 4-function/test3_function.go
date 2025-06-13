package main

import "fmt"

/*
说明: 下面这个例子返回的是单个值
func [函数名]([输入参数] [参数类型], ...) [返回值类型] {
	函数内部
}
*/

// 例子
func foo1 (a string, b int) int {
	fmt.Println("a = ",a)
	fmt.Println("b = ",b)

	c := 100

	return c
}

/*
说明: 下面这个例子返回的是多个值
*/
func foo2(a string, b int) (e int, f int) {
// 假如是两个未设定的匿名返回值 可以写作 unc foo2(a string, b int) (int, int) {
// 如果 两个变量的数据类型相同 也可写成 func foo2(a string, b int) (e, f int) {
	fmt.Println("a = ",a)
	fmt.Println("b = ",b)

	e = 666
	f = 777
	return e, f
}


func main() {
	c := foo1("abc", 555)
	fmt.Println("c = ",c)

	e ,f := foo2("efg", 999)
	fmt.Println("e = ",e)
	fmt.Println("f = ",f)
}