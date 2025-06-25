package main

import (
	"fmt"
	"reflect"
)


type User struct {
	Id int
	Name string
	Age int
}

func (this *User) Call() {
	fmt.Println("user is called")
	fmt.Println(this)
}

func DoFiledAndMethod(input interface{}) {
	// 获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType:", inputType)
	// 获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue:", inputValue)
	// 通过type 获取里面的字段
	// 1.先获取到reflect.Type
	// 2.通过type 获取NumField，进行遍历
	// 3.再通过type 获取到具体的field
	// 4.通过value 获取到具体的field的value
	for i := 0; i < inputType.NumField(); i++ {
			field := inputType.Field(i)
			value := inputValue.Field(i).Interface()	
			fmt.Printf("Field Name: %s, Type: %s, Value: %v\n", field.Name, field.Type, value)
		}
	// 通过type 获取里面的方法

}

func main() {
	u := &User{1, "Tom", 18}
	// u.Call()
	DoFiledAndMethod(u)
}