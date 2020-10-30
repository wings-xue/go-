package main

// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "fmt"

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	fmt.Println(C.srandom(C.uint(i)))
}

func main() {
	fmt.Println(Random())
	Seed(1)
}
