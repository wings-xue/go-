package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 前缀和后缀
	s := "s this company"
	p := "s "
	n := "b "
	r := 3
	fmt.Printf("`%s` have prefix `%s`? %t\n", s, p, strings.HasPrefix(s, p))
	// 包含关系
	fmt.Printf("`%s` contains `%s`? %t\n", s, p, strings.Contains(s, p))

	// 判断索引位置
	fmt.Printf("`%s` location `%s`? %d\n", s, p, strings.Index(s, p))
	// 字符串替换
	fmt.Printf("old: `%s`,new: `%s`, replace: `%s, %s`\n", s, strings.Replace(s, p, n, -1), p, n)
	// 统计字符串出现次数
	fmt.Printf("`%s` occurrence `%s` : %d\n", s, p, strings.Count(s, p))
	// 重复字符串/copy字符串
	fmt.Printf("`%s` repeat '%d': %s\n", s, r, strings.Repeat(s, r))

	// 修改字符串大小写
	fmt.Printf("`%s` Upper `%s`\n", s, strings.ToUpper(s))

	// 修剪字符串
	fmt.Printf("`%s` trim `%s`, rst: %s\n", s, p, strings.Trim(s, p))

	// 分隔字符串
	fmt.Println(strings.Split(s, p))
	// 拼接slice到字符串
	fmt.Println(strings.Join([]string{"h", "e", "l", "l", "o"}, ""))
	// 从字符串读取内容, 感觉蛮有意思的， 写一个I/O流
	reader := strings.NewReader("nihao 中国")

	mybyte := make([]byte, 9)
	_, err := reader.Read(mybyte)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("读取中的文字:`%s`, 剩余的长度%d\n", string(mybyte), reader.Len())

	// 字符串与其他类型转化
	org := "666"
	num, _ := strconv.Atoi(org)
	fmt.Printf("展示实例类型%T\n", num)

	fmt.Printf("类型：%T\n", strconv.Itoa(num))
}
