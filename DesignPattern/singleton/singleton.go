package singleton

/*
	2022-9-26
	定义： 一个类指允许创建一个对象（或者叫实例），那这个类就是一个单例类。这种设计模式就是单例设计模式，简称单例模式
	两种实现：
		饿汉式：简单，可以把问题及早暴露。
*/
// Singleton 饿汉式单例
type Singleton struct{}

var singleton *Singleton

func init() {
	singleton = &Singleton{}
}

// GetInstance 获取实例
func GetInstance() *Singleton {
	return singleton
}
