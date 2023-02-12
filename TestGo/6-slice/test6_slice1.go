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
	// fmt.Println(a[3]) // 不可以访问：panic: runtime error: index out of range [3] with length 2
	fmt.Println(a)
}

func test(a []int) {
	a = append(a, 3) // 虽然在内部append了，但还是不行。
	// a = append(a, 4) // 不发生扩容，下面的更改前面也可以看到。发生扩容后就函数内外不一致了。
	fmt.Println("test:", a, cap(a))
	for i := 0; i < len(a); i++ {
		a[i] *= 2
	}
}
