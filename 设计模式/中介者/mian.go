package mian

// 角色
// 抽象中介者
// 具体中介者
// 抽象同事类
// 具体同事类   中介者调用生成对象在这里实现
// 中介者不打交道， 客户端只与抽象通识类打交道

// type database interface {
// 	add()
// }

// type db struct {
// 	m Mediator
// }

// // 抽象同事类
// func (p db) set(m Mediator) {
// 	p.m = m
// }

// func (p db) add(i int) {
// 	p.m.get(i).add()
// }

// // 具体通识类
// type pgsql struct {
// }

// func (p pgsql) add() {
// 	fmt.Println("pgsql add...")
// }

// type redis struct {
// }

// func (p redis) add() {
// 	fmt.Println("redis add...")
// }

// type Mediator interface {
// 	get(int) database
// }

// // 具体中介者
// type jMediator struct {
// }

// func (j jMediator) get(i int) database {
// 	if i == 0 {
// 		return pgsql{}
// 	} else {
// 		return redis{}
// 	}
// }
