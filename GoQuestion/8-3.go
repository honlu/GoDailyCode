/*
2023-1-3
question:
	写出代码输出，及原因
answer:

*/
package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
这个输出结果决定来自于调度器优先调度哪个G。
从runtime的源码可以看到，当创建一个G时，会优先放入到下一个调度的`runnext`字段上作为下一次优先调度的G。
因此，最先输出的是最后创建的G，也就是9.
(有点难懂)
*/
func test1() {
	runtime.GOMAXPROCS(1) // 设置最大进程数
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

/*
defer 在定义的时候会计算好调用函数的参数，所以会优先输出10、20 两个参数。
然后根据defer定义的顺序倒序执行。
*/
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func test2() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

/*
结果是随机执行。golang 的select在多个`case` 可读的时候会公平的选中一个执行。
*/
func test3() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		// panic(value)
		fmt.Println(value)
	}
}

/*
make 在初始化切片时指定了长度，所以追加数据时会从len(s) 位置开始填充数据。
*/
func test4() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3) // 末尾追加
	fmt.Println(s)
}

func main() {
	// test2()
	// test3()
	test4()
}
