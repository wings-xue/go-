package main

func name(first, second string) string {
	return first + second
}

// 抽象参数
type Personer struct {
	first  string
	second string
}

// 抽象函数
type Namer interface {
	name(*Personer) string
}

// 定义一个结构存储抽象
type funcOption struct {
	f func(*Personer) string
}

// 实现函数接口，具体实现
func (fo funcOption) name(p *Personer) string {

	return fo.f(p)
}

// 抽象函数过程,相当于加了一层代理，用来处理参数
func New(f func(p *Personer) string) funcOption {
	return funcOption{f: f}
}

func First(first string) Namer {

	return New(func(p *Personer) string {
		p.first = first
		return p.first
	})
}
