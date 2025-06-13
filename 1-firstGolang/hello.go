package main // 当前程序包名

import "fmt"
import "time"


\*
import (
	"fmt"
	"time"
)
*\

func main() { // 函数左花括号和函数同行
	// golang表达式建议不加分号
	fmt.Println("hello Go!")

	time.Sleep(1 * time.Second)
}