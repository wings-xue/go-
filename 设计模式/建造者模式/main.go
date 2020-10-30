package main

import "fmt"

// 建造者接口
// 建造者实现
// 产品
// 导演（管理建造流程，传入相应的建造者，用来建造产品）

// 建造者接口
type GirlBuilder interface {
	BuildEyes()
	BuildClothes()
}

// 建造者实现
type LiuYiFei struct {
}

func (l LiuYiFei) BuildEyes() {
	fmt.Println("我是你的女友刘亦菲，我的眼睛是好看的")
}

func (l LiuYiFei) BuildClothes(s string) {
	fmt.Println("我是你的女友刘亦菲，我穿的衣服是好看的")
}

// 建造者实现
type YangMi struct {
}

func (l YangMi) BuildEyes() {
	fmt.Println("我是你的女友杨幂，我的眼睛是灰色的")
}

func (l YangMi) BuildClothes() {
	fmt.Println("我是你的女友杨幂，我穿的衣服是没有穿")
}

// 创建导演
type Director struct {
	build GirlBuilder
}

func (d *Director) set(b GirlBuilder) {
	d.build = b
}

func (d *Director) doBuild() {
	d.build.BuildEyes()
	d.build.BuildClothes()
}

func main() {
	b := YangMi{}
	d := &Director{}
	d.set(b)
	d.doBuild()
}
