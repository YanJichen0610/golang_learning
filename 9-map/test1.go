package	main

import "fmt"

func main() {
	// 声明map1是一种map类型 key是string value是string
	// 代码只是声明了一个 map 变量，但没有进行初始化
	var map1 map[string]string
	if (map1 == nil) {
		fmt.Println("map1 是一个空map")
	}

	map1 = make(map[string]string,10)
	map1["one"] = "java"
	map1["two"] = "c++"
	map1["three"] = "python"

	fmt.Println(map1)

	fmt.Println("---------------------")

	// 第二种声明方式
	map2 := make(map[int]string)
	map2[1] = "java"
	map2[2] = "c++"
	map2[3] = "python"
	fmt.Println(map2)

	fmt.Println("---------------------")

	// 第三种声明方式
	map3 := map[string]string{
		"one": "java",
		"two": "c++",
		"three": "python",
	}
	fmt.Println(map3)
}