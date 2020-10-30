package main

import "fmt"

type Subject interface {
	Do()
}

type S struct {
}

func (s S) Do() {
	fmt.Println(" do")
}

type SProxy struct {
	s S
}

func (s SProxy) Do() {
	fmt.Println("proxy is working")
	s.s.Do()
}

func main() {
	s := SProxy{}
	s.Do()

}
