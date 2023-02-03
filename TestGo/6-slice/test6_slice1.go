package main

import "fmt"

/*
面试询问：
关于slice在传入函数之后的变化，15行代码是否可以正常运行？
*/
func main() {
	a := make([]int, 2, 3)
	a[0] = 1
	a[1] = 2
	fmt.Println(a)
	test(a)
	fmt.Println(a[3]) // 不可以访问：panic: runtime error: index out of range [3] with length 2
}

func test(a []int) {
	a = append(a, 3) // 虽然在内部append了，但还是不行。
	fmt.Println("test:", a, cap(a))
}
