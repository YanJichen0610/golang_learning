package main

// 需求：统计1-100的数字中，哪些是素数

import (
	"fmt"
	"time"
	"sync"
)

var wg sync.WaitGroup

func test1() {
	for num := 2; num <= 1000; num++ {
		var flag = true
		for i := 2; i < num; i++ {
			if num % i == 0 {
				flag = false
				break
			}
		}
		if flag {
			// fmt.Println("素数数字:", num)

		}
	}
}


func test2(num int) {
	var flag = true
	for i := 2; i < num; i++ {
		if num % i == 0 {
			flag = false
			break
		}
	}
	if flag {
		// fmt.Println("素数数字:", num)
	}
	wg.Done()
}


func main() {
	// 单协程测试
	start := time.Now()
	test1()
	end := time.Now()
	fmt.Println("test1执行时间:", end.Sub(start))

	// 多协程测试
	start2 := time.Now()

	for num := 2; num <= 1000; num++ {
		wg.Add(1)
		go test2(num)
	}

	wg.Wait() // 等待所有goroutine完成
	end2 := time.Now()

	fmt.Println("test2执行时间:", end2.Sub(start2))
}