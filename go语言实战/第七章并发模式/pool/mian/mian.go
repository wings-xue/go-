package main

import (
	"gobook/go语言实战/第七章并发模式/pool"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxGoroutines   = 25 // 要使用goroutine的数量
	pooledResources = 2  // 池中资源的数量
)

// dbConnection模拟要共享的资源
type dbConnection struct {
	ID int32
}

// Close实现了io.Closer接口，以便dbConnection可以被池管理。Close用来完成任意资源的释放管理
func (dbConn *dbConnection) Close() error {
	log.Println("Close: Connetion", dbConn.ID)
	return nil
}

// idCounter 用来给每个连接分配一个独一无二的id
var idCounter int32

// createConnection是一个工厂函数，当需要一个新连接时，资源会调用这个函数
func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)
	return &dbConnection{id}, nil

}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	// 创建用来管理连接的池
	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	// 使用池里的连接来完成查询
	for query := 0; query < maxGoroutines; query++ {
		// 每个goroutine需要自己复制一份要查询的副本，不然所有查询会共享同一个查询变量
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()

	// 关闭池
	log.Println("Shutdown Program")
	p.Close()
}

// performQueries用来测试连接池
func performQueries(query int, p *pool.Pool) {
	// 从池里获取连接
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	// 将该连接放回池里
	defer p.Release(conn)

	// 用等待来模拟查询相应
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	// 断言类型，大类型往小类型转
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}
