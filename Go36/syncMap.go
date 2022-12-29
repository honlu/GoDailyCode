package main

import (
	"fmt"
	"reflect"
	"sync"
)

/*
Go 语言自带的字典类型map并不是并发安全的。
sync.Map解决了这个问题，但使用时要注意保证它的键类型和值类型的正确性！
	键的实际类型不能是函数类型、字典类型和切片类型。
	（Go 语言编译器是无法在编译期对它们进行检查的，这些键值的实际类型只能运行期间才能确定，不正确的键值实际类型肯定会引发 panic。）
	必须保证键的类型是可比较的（或者说可判等的），go中不可比较的类型:slic,map,func。
	可以使用reflect.Typeof函数 + Comparable方法类似保证类型的正确性。

两种方案：使用类型断言表达式或者反射操作来保证它们的类型正确性。
保证并发安全字典中的键和值的类型正确性。
第一种方案是，仅仅使用sync.Map在编码时就完全确定键和值的类型，然后利用 Go 语言的编译器帮我们做检查。
第二种方案：利用sync.Map和reflect,接受动态的类型设置，并在程序运行的时候通过反射操作进行检查。
*/

// 第一种实现
type IntStrMap struct { // 自己定义一个key为int型，value为string类型的并发安全map，然后封装它的一些基本方法。
	m sync.Map
}

func (iMap *IntStrMap) Delete(key int) {
	iMap.m.Delete(key)
}

func (iMap *IntStrMap) Load(key int) (value string, ok bool) {
	v, ok := iMap.m.Load(key)
	if v != nil {
		value = v.(string)
	}
	return
}

func (iMap *IntStrMap) LoadOrStore(key int, value string) (actual string, loaded bool) {
	a, loaded := iMap.m.LoadOrStore(key, value)
	actual = a.(string)
	return
}

func (iMap *IntStrMap) Range(f func(key int, value string) bool) {
	f1 := func(key, value interface{}) bool {
		return f(key.(int), value.(string))
	}
	iMap.m.Range(f1)
}

func (iMap *IntStrMap) Store(key int, value string) {
	iMap.m.Store(key, value)
}

// 第二种实现
type ConcurrentMap struct { // 可自定义键类型和值类型的并发安全字典
	m         sync.Map     // 内部使用的并发安全字典
	keyType   reflect.Type // 分别用于保存键类型和值类型。这两个字段的类型都是reflect.Type，我们可称之为反射类型。
	valueType reflect.Type
}

func (cMap *ConcurrentMap) Load(key interface{}) (value interface{}, ok bool) { // 参数key代表了某个键的值。
	if reflect.TypeOf(key) != cMap.keyType { // 把一个接口类型值传入reflect.TypeOf函数，就可以得到与这个值的实际类型对应的反射类型值。
		return
	}
	return cMap.m.Load(key)
}

func (cMap *ConcurrentMap) Store(key, value interface{}) {
	if reflect.TypeOf(key) != cMap.keyType {
		panic(fmt.Errorf("wrong key type: %v", reflect.TypeOf(key)))
	}
	if reflect.TypeOf(value) != cMap.valueType {
		panic(fmt.Errorf("wrong value type: %v", reflect.TypeOf(value)))
	}
	cMap.m.Store(key, value)
}
