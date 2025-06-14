package main

import "fmt"


func printMap(map_data map[string]string)  {
	// map_data是引用传递 指向同一个空间
	for key, value := range map_data {
		fmt.Println(key, value)
	}
}

func modifyMap(map_data map[string]string)  {
	// 因为是引用传递 指向同一个空间
	// 所以修改会影响到map_data
	map_data["USA"] = "DC"
}

func main() {
	map1 := make(map[string]string)
	// map中添加信息
	map1["China"] = "beijing"
	map1["USA"] = "newyork"
	map1["Japan"] = "tokyo"

	// 遍历
	printMap(map1)

	fmt.Println("---------------------")

	// 删除
	delete(map1, "China")
	// 再次遍历
	printMap(map1)

	fmt.Println("---------------------")
	// 修改
	map1["USA"] = "wasington"
	// 再次遍历
	printMap(map1)

	// 使用方法修改
	modifyMap(map1)
	fmt.Println("---------------------")
	// 再次遍历
	printMap(map1)
}
