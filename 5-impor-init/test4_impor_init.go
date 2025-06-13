/*
在 Go 中，程序的执行从 main 包开始。

Go 运行时会从 main 包中的 main 函数作为程序入口，
但在这之前会按照以下顺序进行初始化操作：

1. 通过 import 导入所依赖的包。每个被导入的包只会初始化一次。
2. 对每个包，会先按顺序执行该包中所有的变量声明（包括全局变量的初始化）。
3. 然后调用该包中定义的 init() 函数（如果有多个 init()，按源文件顺序调用）。
4. 所有依赖包初始化完毕后，最后才会执行 main 函数。

init 函数不需要手动调用，也不能被其他函数调用，主要用于包级初始化工作。
每个包可以包含多个 init 函数，甚至每个文件都可以有自己的 init。
*/
package main
// import "fmt"

import (
	"GolangStudy/5-impor-init/lib1"
	"GolangStudy/5-impor-init/lib2"
)
/*
1. 匿名导入说明
仅执行包的 init() 函数，而不会将包的标识符引入到当前作用域。
这种方式常用于导入包的初始化逻辑，而不需要使用包中的标识符。
匿名导入使用方法如下：
import _ "package/path"
其中，_ 是一个空白标识符，用于表示导入包但不使用其中的标识符。
例如：
import _ "fmt" // 导入 fmt 包，但不使用其中的标识符
*/

/*
2. 别名导包
import (
    别名 "包路径"
)
*/

func main () {
	lib1.Lib1Test()
	lib2.Lib2Test()
}