package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner在给给定的超时时间执行一组任务，并且在操作系统发送终端信号时结束这些任务
type Runner struct {
	// interrupt通道报告从操作系统发送的信号
	interrupt chan os.Signal

	// complete通道报告处理任务已经完成
	complete chan error

	// timeout 报告处理任务已经超时
	// 接受time.time 类型的数据
	timeout <-chan time.Time

	// tasks持有一组以索引顺序依次执行的函数
	tasks []func(int)
}

// 定义err
// ErrTimeout 会在执行任务超时时返回
var ErrTimeout = errors.New("received timeout")

// ErrInterrupt 会在收到操作系统事件时返回
var ErrInterrupt = errors.New("received interrupt")

// New返回一个新的准备使用的Runner
func New(d time.Duration) *Runner {
	return &Runner{
		// interrupt被初始化缓冲区容量为1的通道.这可以保证通道至少能接受一个来自语言运行时的os.Signal值，确保语言运行时发送这个事件不会被阻塞。
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		// 在等待duration去释放然后发送当前时间到channel
		timeout: time.After((d)),
	}
}

// Add将一个任务附加到Runner上。这个任务是一个接受一个int类型的ID作为参数的函数
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start执行所有任务，并监视隧道事件
func (r *Runner) Start() error {
	// 我们希望接受所有终端信号
	// os.Interrupt所有的信息转发给r.interrupt
	signal.Notify(r.interrupt, os.Interrupt)

	// 用不同的goroutine执行不同的任务
	go func() {
		r.complete <- r.run()
	}()

	select {
	// 如果键盘输入ctrl+c返回err
	case err := <-r.complete:
		return err
	// 当任务处理程序运行超时时发的信号
	case <-r.timeout:
		return ErrTimeout
	}

}

// 运行的时候杀死
func (r *Runner) run() error {
	for id, task := range r.tasks {
		// 检查是否有ctrl+c
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		task(id)
	}
	return nil
}

func (r *Runner) gotInterrupt() bool {
	select {
	// 当终端事件被触发时发出的信号

	case <-r.interrupt:
		// 停止接受后续任何信息
		signal.Stop(r.interrupt)
		return true
	// 继续正常运行
	default:
		return false
	}

}
