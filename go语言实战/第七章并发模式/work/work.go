package work

import "sync"

// Worker必须满足接口类型才能工作
type Worker interface {
	Task()
}

// Pool提供一个goroutine池，这个池可以完成任何已提交的Worker任务
type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

func New(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}
	p.wg.Add(maxGoroutines)

	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}
	return &p
}

// Run提交工作到工作池
func (p *Pool) Run(w Worker) {
	p.work <- w
}

func (p *Pool) Shutdonw() {
	close(p.work)
	p.wg.Wait()
}
