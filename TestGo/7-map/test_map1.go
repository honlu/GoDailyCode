/*
面试题：实现map中value可以多个类型
value为interface{}就可以实现。
不过interface{}在使用时也要转换一下格式。
可以通过类型断言、类型选择来使用，如果没有学过反射。
也可以使用反射来使用。
反射相关文章：
	https://zhuanlan.zhihu.com/p/371740353 反射的基础知识
	https://zhuanlan.zhihu.com/p/372418957 使用反射修改值

*/
package main

import "fmt"

func main() {
	// mapStr := make(map[string]string)    // 只能储存string类型的value
	// mapInt := make(map[string]int) // 只能保存int类型的value

	hash := make(map[string]interface{}) //可以存放string的key，任意类型（比如string,int等）的value
	hash["id"] = 1
	hash["plan_id"] = "第一个"
	//fmt.Println(hash["id"]+1)  会报错(mismatched types interface {} and int)
	fmt.Println(hash["id"].(int) + 1)

}
