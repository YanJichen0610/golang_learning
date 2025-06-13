package main

import "fat"

// const 来定义枚举类型 如下所示
const (
	BEIJING = 0
	SHANGHAI = 1
	GUANGZHOU = 2
)

// 可以在const()中追加关键字iota, 每一行的iota会进行累,第一行的iota为0
// 注意iota只能配合const()进行使用
/*
const (
	BEIJING = iota	
	SHANGHAI
	GUANGZHOU
)
*/

func main() {
	// 常量只读属性(不可赋值)	
	const length int = 10 
}