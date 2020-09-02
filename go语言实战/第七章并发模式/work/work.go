package work

import (
	"fmt"
	"sync"
)

// Worker必须满足接口类型才能工作
type Worker interface {
	Task()
}

// Pool提供一个goroutine池，这个池可以完成任何已提交的Worker任务
type Pool struct {
	Work chan Worker
	wg   sync.WaitGroup
}

func New(maxGoroutines int) *Pool {
	p := Pool{
		Work: make(chan Worker),
	}
	p.wg.Add(maxGoroutines)

	for i := 0; i < maxGoroutines; i++ {
		fmt.Printf("PROCESS-%d \n", i)
		// 如果go func 里面直接调用i，调用的是i的引用
		go func(i int) {

			for w := range p.Work {
				w.Task()

			}
			// 这里永远不会关闭，所以这套代码存在问题
			fmt.Printf("channel 读取完毕.....")
			p.wg.Done()
		}(i)
	}
	return &p
}

// Run提交工作到工作池
func (p *Pool) Run(w Worker) {
	p.Work <- w
}

func (p *Pool) Shutdonw() {
	close(p.Work)
	p.wg.Wait()
}
