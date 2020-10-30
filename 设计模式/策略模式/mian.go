package mian

import "fmt"

// 模板和具体执行分离

type Reader struct {
	strategy readliner
}

func (r *Reader) set(s readliner) {
	r.strategy = s
}

func (r *Reader) do() {
	r.strategy.do()
}

type readliner interface {
	do()
}

type paper struct {
}

func (p paper) do() {
	fmt.Println("我正在读")
}
