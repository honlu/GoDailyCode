package main

import "fmt"

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
