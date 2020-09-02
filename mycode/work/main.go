package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// 1. for channal如果不close会一直阻塞

func read(c chan int, wc sync.WaitGroup) {

	for i := range c {
		log.Printf("channel value is %d", i)
		defer wc.Done()
	}
}

func add(c chan int) {
	c <- 1
}

// right code
func chan2() {
	messages := make(chan int)
	var wg sync.WaitGroup
	// you can also add these one at
	// a time if you need to
	wg.Add(3)
	go func() {
		defer wg.Done()
		// time.Sleep(time.Second * 3)
		messages <- 1
	}()
	go func() {
		defer wg.Done()
		// time.Sleep(time.Second * 2)
		messages <- 2
	}()
	go func() {
		defer wg.Done()
		// time.Sleep(time.Second * 1)
		messages <- 3
	}()
	go func() {
		for i := range messages {
			fmt.Println(i)
		}
	}()
	wg.Wait()
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

func chan3() {
	var wc sync.WaitGroup
	wc.Add(3)

	ch := make(chan int)
	for i := 0; i < 3; i++ {
		go func() {
			ch <- 1
			wc.Done()
		}()

	}
	go func() {
		// for range 会一直开启
		for i := range ch {
			fmt.Println(i)
		}
		// 这里是不会打印的
		fmt.Println("READ CLOSE")
	}()
	time.Sleep(time.Second * time.Duration(10))
	wc.Wait()

}
func main() {
	// forchannel()
	// chan2()
	chan3()
}
