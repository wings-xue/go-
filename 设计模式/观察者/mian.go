package main

import "fmt"

// 观察者模式
// 角色
// 目标
// 具体目标   目标里面存放具体观察者，目标更新，调用观察者的函数执行

// 观察者
// 具体观察者

type subject struct {
	o       []Observer
	content string
}

type subjecter interface {
	update()
	registry(Observer)
	change(string)
}

func (s *subject) update() {
	for _, obj := range s.o {
		obj.Update()
	}
}

func (s *subject) registry(o Observer) {
	s.o = append(s.o, o)
}

func (s *subject) change(str string) {
	s.content = str
	s.update()
}

type Observer interface {
	Update()
}

type drawer struct {
}

func (d drawer) Update() {
	fmt.Println("draw update")
}
