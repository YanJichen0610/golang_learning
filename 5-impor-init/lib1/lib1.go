// 一般包名与文件名统一 方便查找
package lib1
import "fmt"

// 注意函数名首字母大写 才可以被外部调用
// 函数名首字母小写 只能在本包中调用
func Lib1Test () {
	fmt.Println("lib1 test")
}

func init() {
	fmt.Println("lib1 init")
}