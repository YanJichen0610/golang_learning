package main

import "fmt"

func main() {
	// pair<statictype:string, value:"aceld">
	var a string
	a = "aceld"

	// 这里指针指向'a'
	// pair<type:string, value:"aceld">
	var allType interface{}
	allType = a

	str, _ := allType.(string)
	fmt.Println(str)
}