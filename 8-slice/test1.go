package main
import "fmt"

func main() {
	// 定义数组基本语法
	// 写法1：
	// var 数组名 [数组长度]数据类型
	var arr1 [10]int

	// 写法2：
	// var 数组名 := [数组长度]数据类型{值1, 值2, ...}
	var arr2 = [10]int{1, 2, 3, 4, 5}

	// 遍历数组
	for i := 0; i < len(arr1); i++ {
		// 这里打印初始值为0
		fmt.Println(arr1[i])
	}

	fmt.Println("----------------")

	// 遍历数组
	for index, value := range arr2 {
		fmt.Printf("index = %d, value = %d\n", index, value)
	}

	// 查看数组的数据类型
	fmt.Printf("arr1的数据类型为：%T\n", arr1)
	fmt.Printf("arr2的数据类型为：%T\n", arr2)
}