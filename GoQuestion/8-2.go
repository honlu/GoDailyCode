/*
2023-1-3
question:
	下面代码的问题
answer:
*/
package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu // 这里有问题： `for ... range` 语法中，`stu` 变量会被复用。由于是地址传递，每次循环会将集合中的值复制给这个变量，因此，会导致最后`m`中的`map`中储存的都是`stus`最后一个`student`的值。
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}

// 正确保存方式
func pase_student2() {
	m := make(map[string]student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = stu // 这里有问题： `for ... range` 语法中，`stu` 变量会被复用。由于是地址传递，每次循环会将集合中的值复制给这个变量，因此，会导致最后`m`中的`map`中储存的都是`stus`最后一个`student`的值。
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}

func main() {
	// pase_student()
	pase_student2()
}
