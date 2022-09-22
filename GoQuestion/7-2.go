package main

/*
2022-9-22
question:
	类型断言、switch匹配问题
answer:
	golang中有规定，switch type的case T1，类型列表只能有一个，那么当v := m.(type)中的v的类型就是T1类型。
	如果是case T1, T2，类型列表中有多个，那v的类型还是对应接口的类型即interface{}，也就是m的类型。
	所以这里msg的类型还是interface{}，所以他没有Name这个字段，编译阶段就会报错。
*/
import "fmt"

type student struct {
	Name string
}

func zhoujielun(v interface{}) {
	switch msg := v.(type) {
	case *student: // type of i is *student
		fmt.Println(msg.Name)
	case student: // type of i is student
		fmt.Println("a", msg.Name)
	// 下面是错误写法，编译就报错
	// case *student, student: // type of i is type of v (interface{})
	// 	fmt.Println("a", msg.Name)
	// }

}

func main() {
	s := student{Name: "zhoujielun"}
	zhoujielun(s)  // 打印： a zhoujielun
}
