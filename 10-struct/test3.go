package main

import "fmt"

type Human struct {
	Name string
	Sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

type SuperMan struct { // 修复错误：s truct -> struct
	Human // 继承了 Human 的属性和方法
	Level int
}

// 重新定义父类方法
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}

// 新增子类方法
func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}

func main() {
	h := Human{
		Name: "zhang3",
		Sex:  "男",
	}
	h.Eat()
	h.Walk()

	fmt.Println("---------------------")

	// 定义子类对象
	// s := SuperMan{Human: Human{Name: "li4",Sex:  "男",},Level: 88,}
	var s SuperMan
	s.Name = "li4"
	s.Sex = "男"
	s.Level = 88

	s.Fly()   // 子类方法
	s.Walk()  // 继承的方法
	s.Eat()   // 重写的方法
}
