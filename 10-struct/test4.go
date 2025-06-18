package main


import (
	"fmt"
)


// interface 本质是一个指针
type AnimalIF interface {
	Sleep()
	GetColor() string // 获取动物颜色
	GetType() string // 获取动物种类
}

// 具体的种类
// 重写接口方法
type cat struct {
	color string
}

func (this *cat) Sleep() {
	fmt.Println("Cat is sleeping")
}

func (this *cat) GetColor() string {
	return this.color
}

func (this *cat) GetType() string {
	return "cat"
}


// 增加一个类
type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is sleeping")
}


func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string{
	return "dog"
}

func showAnimal(animal AnimalIF) {
	animal.Sleep()
}

func main() {
	// 定义一个 AnimalIF 接口变量，初始化为 cat
	var animal AnimalIF

	animal = &cat{color: "black"} // 使用短变量声明
	fmt.Println(animal.GetType())
	animal.Sleep()
	showAnimal(animal)
	fmt.Println("------------------------------")

	// 将 animal 赋值为另一个具体类型（Dog），注意用 = 而不是 :=
	animal = &Dog{color: "white"}
	fmt.Println(animal.GetType())
	animal.Sleep()
	showAnimal(animal)

	dog := Dog{color: "white"}
	cat := cat{color: "black"}
	
	showAnimal(&dog)
	showAnimal(&cat)
}
