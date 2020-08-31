package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func coolchan() {
	// 三种chan
	c1 := make(chan int)
	c2 := make(<-chan int)
	c3 := make(chan<- int)

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
}

// time.After测试
func coolAfter() {
	go func() {
		time.Sleep(time.Second * 2)
	}()

	select {
	case <-time.After(time.Second * 1):
		log.Println("time out")
	}
}

func coolsignal() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool)

	signal.Notify(sigs, syscall.SIGABRT)

	go func() {
		fmt.Println("wait for signal")
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()

	<-done

}

type runner struct {
	// 用来接受信号
	interrupt chan os.Signal
	// 用来接受运行结果中出现的错误
	complete chan error
	// 用来接受是否超时
	timeout <-chan time.Time
	// 任务队列
	task []func(int)
}

func New(c time.Duration) *runner {
	interrupt := make(chan os.Signal, 1)
	complete := make(chan error)
	fmt.Println(c)
	return &runner{
		interrupt: interrupt,
		complete:  complete,
		timeout:   time.After(c),
	}
}

var InterruptErr = errors.New("InterruptErr")
var TimeoutErr = errors.New("TimeoutErr")

func (r *runner) start() error {
	log.Println("Process start...")
	// 绑定信号
	signal.Notify(r.interrupt, os.Interrupt)

	// 执行任务
	go func() {

		r.complete <- r.run()
	}()

	select {
	// 中断信号获取
	case <-r.complete:

		return InterruptErr
	case <-r.timeout:
		return TimeoutErr
	}
}

// 执行单个任务
func (r *runner) run() error {

	for id, task := range r.task {
		// 获取 ctrl+c信号
		if r.getinterrupt() {
			return InterruptErr
		}
		task(id)
	}
	return nil
}

// 获取ctrl+c信号
func (r *runner) getinterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}

}

func (r *runner) add(tasks func(int)) {
	r.task = append(r.task, tasks)
}

func printID(id int) {
	time.Sleep(time.Duration(id) * time.Second)
	log.Printf("Processor - Task #%d", id)
}

func main() {
	// coolchan()
	// // 测试time.After使用
	// coolAfter()

	// coolsignal()
	taskNum := 5
	runner := New(time.Duration(taskNum) * time.Second)
	for i := 0; i < taskNum; i++ {
		runner.add(printID)
	}
	if err := runner.start(); err != nil {
		switch err {
		case InterruptErr:
			os.Exit(1)
		case TimeoutErr:
			os.Exit(2)
		}

	}

}
