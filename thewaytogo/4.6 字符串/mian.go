package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// byte 等同于int8，常用来处理ascii字符
	// rune 等同于int32,常用来处理unicode或utf-8字符
	var str = "hello 你好"

	// golang中string底层是通过byte数组实现的。中文字符在unicode下占2个字节，在utf-8编码下占3个字节(byte)，而golang默认编码正好是utf-8。
	// str 获取字节数
	fmt.Println("len(str):", len(str))

	// golang中的unicode/utf8包提供了用utf-8获取长度的方法
	// 获取utf8 单位
	fmt.Println("RuneCountInString:", utf8.RuneCountInString(str))

	//通过rune类型处理unicode字符
	fmt.Println("rune:", len([]rune(str)))

	s1 := "asSASA ddd dsjkdsjs dk"
	s2 := "asSASA ddd dsjkdsjsこん dk"

	fmt.Printf("%s 字节数:%d \n", s1, len(s1))
	fmt.Printf("%s utf8数:%d \n", s1, utf8.RuneCountInString(s1))

	fmt.Printf("%s 字节数:%d \n", s2, len(s2))
	fmt.Printf("%s utf8数:%d \n", s2, utf8.RuneCountInString(s2))
}
