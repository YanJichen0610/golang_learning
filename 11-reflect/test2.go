package main

import "fmt"

type Reader interface{
	ReadBook()
}

type Writer interface{
	WriteBook()
}

type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("read a book")
}

func (this *Book) WriteBook() {
	fmt.Println("write a book")
}

func main() {
	// b: pair<type:Book, value:book{}地址>
	b := &Book{}

	// r: pair<type:, value:>
	var r Reader
	// pair<type:Reader, value:book{}地址>
	r = b

	r.ReadBook()

	var w Writer
	// r: pair<type:book, value:book{}地址>
	w = r.(Writer) // 此处断言成功 因为 r的type是Book 而Book实现了Reader和Writer的接口
	w.WriteBook()
	
}