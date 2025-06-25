package main

/**
在主线程中（也可以理解为进程），开启一个goroutine，该协程每隔50毫秒输出“hello world”
在主线程中也每隔50毫秒输出“hello golang”，输出10次后，退出程序
要求主线程和goroutine同时执行
**/
import (
	"fmt"
	"time"
	"sync"
)

// 定义一个全局变量wg
var wg sync.WaitGroup

func test1() {
	for i := 0; i < 15; i ++ {
		fmt.Println("hello golang test1", i)
		time.Sleep(time.Millisecond * 50)
	}
	// 协程执行完毕 计数器减一
	wg.Done()
}

func test2() {
	for i := 0; i < 15; i ++ {
		fmt.Println("hello golang test2", i)
		time.Sleep(time.Millisecond * 50)
	}
	// 协程执行完毕 计数器减一
	wg.Done()
}

func main() {
	// 假如直接调用则按顺序执行
	// test()
	// 开启一个协程 并执行test函数 但是这种情况下主线程结束 协程也会跟着结束
	// go test()

	// 为了避免这种情况，可以在这里加上协程计数器
	wg.Add(1)
	go test1()

	wg.Add(1)
	go test2()

	// 主线程执行
	count := 10
	for i := 0; i < count; i++ {
		fmt.Println("hello golang main", i)
		time.Sleep(time.Millisecond * 100)
	}

	// 等待协程执行完毕...
	wg.Wait()
	fmt.Println("主线程退出...")
}