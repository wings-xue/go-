package main

import "fmt"

// 抽象中间层，基于底层重新设计接口，使得面向的接口减少

// 混乱不堪的调用
func chaosBuild() {
	k := Keyboard{}
	k.plug()
	m := Mouse{}
	m.plug()
	// other
}

// 中间调用者
func MidBuild() {
	k := Keyboard{}
	k.plug()
	m := Mouse{}
	m.plug()
}

// 外观模式调用
func facadeBuild() {
	MidBuild()
	// other
}

// 底层实现者
type Devicer interface {
	plug()
}

type Keyboard struct {
}

func (k Keyboard) plug() {
	fmt.Println("键盘已经接入")
}

type Mouse struct {
}

func (k Mouse) plug() {
	fmt.Println("鼠标已经接入")
}

func main() {
	chaosBuild()
	facadeBuild()
}
