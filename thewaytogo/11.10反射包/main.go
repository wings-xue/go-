package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func main() {
	x := 6.3824
	fmt.Println(reflect.TypeOf(x))
	v := reflect.ValueOf(&x) // &获取指针位置  *获取值
	fmt.Println(v.CanSet())
	v = v.Elem()
	fmt.Println(v.CanSet())
	v.SetFloat(6.824)
	fmt.Println(x)

	fmt.Println("*********this is struct print *********")

	h := Happy{why: "maybe no heard"}
	fmt.Println(reflect.TypeOf(h))
	v = reflect.ValueOf(h)
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, v.Field(i))
	}
	result := v.Method(0).Call(nil)
	fmt.Println(result)

	print(price(2.123), Week{day: "sunday"}, "adfadf")
}

type Happy struct {
	why string
}

func (h Happy) Strings() string {
	return h.why
}

type price float64

func (p price) Strings() string {
	return strconv.FormatFloat(float64(p), 'E', -1, 64)
}

type Week struct {
	day string
}

func (w Week) Strings() string {
	return w.day
}

func print(args ...interface{}) {
	for _, arg := range args {
		switch a := arg.(type) { // type switch
		case Week:
			os.Stdout.WriteString(a.Strings())
		case int:
			os.Stdout.WriteString(strconv.Itoa(a))
		case string:
			os.Stdout.WriteString(a)
		case price:
			os.Stdout.WriteString(a.Strings())
		// more types
		default:
			os.Stdout.WriteString("???")
		}
	}
}
