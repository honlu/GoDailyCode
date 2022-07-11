package singleton

import "sync"

/*
sync.Once 是 Golang package 中使方法只执行一次的对象实现，作用与 init 函数类似。但也有所不同。

init 函数是在文件包首次被加载的时候执行，且只执行一次
sync.Onc 是在代码运行中需要的时候执行，且只执行一次
当一个函数不希望程序在一开始的时候就被执行的时候，我们可以使用 sync.Once 。
*/
var (
	lazySingleton *Singleton
	once          = &sync.Once{} // 实现懒汉式不可避免需要加锁
)

// GetLazyInstance 懒汉式
func GetLazyInstance() *Singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &Singleton{}
		})
	}
	return lazySingleton
}
