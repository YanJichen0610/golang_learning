package main

import "fmt"

func main() {
	// 声明slice1是一个切片，并且初始化，默认值是1,2,3，长度len是3
	slice1 := []int{1, 2, 3 }
	fmt.Printf("slice1的长度为：%d\n", len(slice1))
	fmt.Printf("slice = %v\n", slice1)

	fmt.Printf("--------------------------\n")
	// 声明slice是一个切片，但是没有给slice开辟空间
	// var slice []int

	// 声明slice2是一个切片，但是没有初始化，默认值是0，长度len是0
	slice2 := []int{}
	fmt.Printf("slice2的长度为：%d\n", len(slice2))
	fmt.Printf("slice2 = %v\n", slice2)

	fmt.Printf("--------------------------\n")
	
	// 假如要给没有初始化的slice2添加元素，需要开辟空间
	// 这里用mack()方法开辟了三个空间，每个空间的默认值是0
	slice2 = make([]int, 3)
	fmt.Printf("slice2的长度为：%d\n", len(slice2))
	fmt.Printf("slice2 = %v\n", slice2)

	fmt.Printf("--------------------------\n")
	// 声明方法，直接通过slice3来开辟空间
	slice3 := make([]int, 3)
	fmt.Printf("slice3的长度为：%d\n", len(slice3))
	fmt.Printf("slice3 = %v\n", slice3)


}

