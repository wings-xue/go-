package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func Singleton() {
	once.Do(func() {
		fmt.Println("我正在被创建")
	})
}

func main() {
	for i := 0; i < 100; i++ {
		// 这里仅仅执行一次函数
		Singleton()
	}
}
