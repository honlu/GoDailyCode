package main

/*
2022-9-22
question:
	代码问题
answer：
	map初始化，new和make的区别
*/
type Param map[string]interface{}

type Show struct {
	Param
}

func main() {
	s := new(Show)
	// 错误写法
	// s.Param["RMB"] = 10000 // new不能初始化Show结构体内的Param属性,注意：nil map不能直接赋值
	// 正确写法
	s.Param = make(Param)  // map必须先初始化,make可以，new不能
	s.Param["RMB"] = 10000 // 才能赋值
}
