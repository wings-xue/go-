package main

import (
	"fmt"
	"gobook/go语言实战/第七章并发模式/work"
	"log"
	"sync"
	"time"
)

var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

// namePrinter使用特定方式打印名字
type namePrinter struct {
	name string
}

// Task实现Worker接口
func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

func main() {
	// 使用两个goroutine来创建工作池
	p := work.New(2)

	var wg sync.WaitGroup

	wg.Add(10 * len(names))
	// 这里的数量应该和wg add的数量相同，否则会报错
	for i := 0; i < 10; i++ {
		// 迭代names切片
		for _, name := range names {
			// 创建一个namePrinter并提供指定的名字
			np := namePrinter{
				name: name,
			}

			go func() {
				// 将任务提交执行。当Run返回时我们就知道任务已经处理完成
				p.Run(&np)
				wg.Done()
			}()

		}
	}

	wg.Wait()
	fmt.Println("推送完毕")

	// 让工作池停止工作，等待所有现有的工作完成
	p.Shutdonw()

}
