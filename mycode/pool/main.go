package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"sync"
	"time"
)

// 线程池

// 锁使用方法

type SafeCounter struct {
	mux sync.Mutex
	v   map[string]int
}

func (s *SafeCounter) Inc(key string) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.v[key]++
}

func (s *SafeCounter) Value(key string) int {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.v[key]
}

func Run() {
	// nil map 不可以赋值
	s := &SafeCounter{v: make(map[string]int)}
	var wc sync.WaitGroup
	wc.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			s.Inc("bob")
			wc.Done()
		}()
	}
	wc.Wait()
	fmt.Println(s.Value("bob"))

}

// 线程池代码
type Pool struct {
	mtx         sync.Mutex
	factory     func() (io.Closer, error)
	resources   chan io.Closer
	close       bool
	size        int
	currentSize int
}

// New
func New(size int, factory func() (io.Closer, error)) *Pool {
	resources := make(chan io.Closer, size)
	return &Pool{
		size:      size,
		factory:   factory,
		resources: resources,
	}
}

// ErrPoolClosed表示请求(Acquire)了一个已经关闭的池
var ErrPoolSizeover = errors.New("SIZE OVER.")

// Acquire
func (p *Pool) Acquire() (io.Closer, error) {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	p.currentSize++
	select {
	case resource := <-p.resources:
		return resource, nil
	default:
		// 自动生成一个
		r, _ := p.factory()
		if p.currentSize > p.size {
			log.Printf("超出池大小,等待size:%d, currentSize:%d", p.size, p.currentSize)
			return r, ErrPoolSizeover
		}

		return r, nil
	}

}

// Release
func (p *Pool) Release(r io.Closer) {
	// 放入队列
	p.mtx.Lock()
	defer p.mtx.Unlock()
	select {
	case p.resources <- r:
		log.Println("资源已经放入资源池汇总")
	default:
		log.Println("资源池已经满了，丢弃")
	}
}

// Close
func (p *Pool) Close() {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	log.Println("close pools")
	close(p.resources)

	for r := range p.resources {
		r.Close()
	}

}

type pgsqler struct {
}

func (conn pgsqler) Close() error {
	return nil
}

func (conn pgsqler) Query() {
	log.Println("开始执行查询函数")
	time.Sleep(time.Duration(3) * time.Second)
}

func createConn() (io.Closer, error) {
	fmt.Println("create a new conn")
	return pgsqler{}, nil
}
func anyPostive(values ...int) bool {
	return false
}

func main() {
	// Run()
	// var wg sync.WaitGroup
	// loopcount := 20
	// size := 10
	// wg.Add(loopcount)
	// pool := New(size, createConn)
	// for i := 0; i < loopcount; i++ {
	// 	go func() {
	// 		conn, _ := pool.Acquire()
	// 		pgcon := conn.(pgsqler)
	// 		pgcon.Query()
	// 		pool.Release(pgcon)

	// 		wg.Done()
	// 	}()
	// }
	// wg.Wait()
	// pool.Close()
	// 测试如果函数的传参是slice是否可以直接调用
	anyPostive()
}
