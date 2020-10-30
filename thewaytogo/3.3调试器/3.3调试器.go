package main

import "fmt"

type Fuck struct {
	name string
}

func main() {
	f := Fuck{
		name: "every body",
	}
	// 打印实例值
	fmt.Printf("展示实例的值:%v\n", f)

	// 打印实例类型名称和值
	fmt.Printf("展示类型名称和值%+v\n", f)

	// 打印实例类型
	fmt.Printf("展示实例类型%T", f)

}
