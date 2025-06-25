// 通过for循环开启多个goroutine
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup


func test(i int) {
	fmt.Println("goroutine编号:", i)
	wg.Done()
}



func main() {
	for i := 0; i < 10; i ++ {
		wg.Add(1)
		go test(i)
	}
	wg.Wait()
}