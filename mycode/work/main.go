package main

import (
	"log"
	"sync"
)

func read(c chan int, wc sync.WaitGroup) {

	for i := range c {
		log.Printf("channel value is %d", i)
		defer wc.Done()
	}
}

func add(c chan int) {
	c <- 1
}

func forchannel() {
	// 做一个函数读取chan
	var wc sync.WaitGroup
	c := make(chan int)
	wc.Add(2)
	go read(c, wc)

	go add(c)
	go add(c)
	wc.Wait()

}

func main() {
	forchannel()
}
