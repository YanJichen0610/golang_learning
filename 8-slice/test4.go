// 切片的追加
package main

import "fmt"

func main() {
	numbers := make([]int, 3, 5)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	fmt.Printf("--------------------------\n")

	// 向numbers切片追加一个元素1
	numbers = append(numbers, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// 向一个容量已满的切片追加元素，会导致切片的重新分配
	// 容量会自动增加，增加的容量是当前容量的一倍
	fmt.Printf("--------------------------\n")
	numbers = append(numbers, 2)
	numbers = append(numbers, 3)
	numbers = append(numbers, 4)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)
}