package main

import "fmt"

// 简单工厂模式

// 产品抽象
type Producter interface {
	name()
}

// TV产品实现
type TV struct {
}

func (t TV) name() {
	fmt.Println("i am TV")
}

// VCD产品实现
type VCD struct {
}

func (t VCD) name() {
	fmt.Println("i am VCD")
}

// 工厂
func factory(key string) Producter {
	switch {
	case key == "vcd":
		return VCD{}
	case key == "tv":
		return TV{}
	}
	return nil
}

func main() {
	f := factory("vcd")
	f.name()
}
