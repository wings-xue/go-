package main

import "fmt"

// 通过工厂子类实现产品，而不是直接实现产品

// 产品
type Personer interface {
	behavior()
}

type JianRen struct {
}

// 产品
func (j JianRen) behavior() {
	fmt.Println("我是一个贱人，我正在犯贱")
}

type PuTongRen struct {
}

// 产品
func (j PuTongRen) behavior() {
	fmt.Println("我是一个普通人，我正在追求利益")
}

type ShuXingRen struct {
}

func (j ShuXingRen) behavior() {
	fmt.Println("我是一个树形人，我正在犯傻然后反杀")
}

type factory interface {
	get() Personer
}

// 工厂子类
type PersonFactory struct {
}

func (p PersonFactory) get() Personer {
	return ShuXingRen{}
}

// 工厂子类
type AnimalFactory struct {
}

func (p AnimalFactory) get() Personer {
	return JianRen{}
}

func main() {
	a := AnimalFactory{}
	a.get().behavior()

	p := PersonFactory{}
	p.get().behavior()
}
