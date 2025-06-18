package main

import "fmt"

// 如果类名首字母大写 表示其他包也可以访问
// 如果类名首字母小写 表示其他包不能访问
type Hero struct {
	// 如果属性首字母大写 表示该属性对外能够访问
	// 如果属性首字母小写 表示该属性只能够类的内部访问
	Name  string
	Ad    int
	Level int
}

func (this Hero) Show() {
	fmt.Println("Name =", this.Name)
	fmt.Println("Ad =", this.Ad)
	fmt.Println("Level =", this.Level)
}

func (this Hero) GetName() string {
	fmt.Println(this.Name)
	return this.Name
}

// 测试方法是否有效
// 发现方法无效 因为是值传递 所以不会改变原来的值
func (this Hero) SetName(newName string) {
	this.Name = newName
}

// 此时不是值传递 而是执政传递 所以调用方法会改变值
// 通常定义类都是要传递地址的 就是要加*号
func (this *Hero) SetNameNew(newName string) {
	this.Name = newName
}

func main() {
	hero := Hero{Name: "zhang3", Ad: 100, Level: 1}
	hero.Show()

	fmt.Println("---------------------")

	hero.SetName("li4")
	hero.Show()

	fmt.Println("---------------------")
	hero.SetNameNew("li5")
	hero.Show()
}
