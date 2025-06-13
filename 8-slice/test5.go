// 切片的截取
package main

import "fmt"

func main() {
	var slice1 = []int{1,2,3} // len = 3, cap = 3
	slice2 := slice1[0:2] // slice = {1, 2} len = 2, cap = 3
	fmt.Printf("slice1 = %v\n", slice1)
	fmt.Printf("slice2 = %v\n", slice2)

	fmt.Printf("--------------------------\n")
	
	// 进行值的替换
	// 此时slice1与slice2的值都会发生改变，因为指向相同的地址
	slice2[0] = 100
	fmt.Printf("slice1 = %v\n", slice1)
	fmt.Printf("slice2 = %v\n", slice2)

	fmt.Printf("--------------------------\n")

	// copy 可以将底层数组信息进行拷贝
	// 也就是让slice3与slice1指向的地址不同
	slice3 := make([]int, 3)
	copy(slice3, slice1)
	fmt.Printf("slice3 = %v\n", slice3)

	slice3[0] = 1000
	fmt.Printf("slice1 = %v\n", slice1)
	fmt.Printf("slice2 = %v\n", slice2)
	fmt.Printf("slice3 = %v\n", slice3)
}