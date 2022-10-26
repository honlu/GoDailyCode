/*
2022-9-23
question:
	sync.Map的用法，并发map.
	判断下面代码的输出
answer:

*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	m.Store("address", map[string]string{"province": "江苏", "city": "南京"})
	v, _ := m.Load("address") // 注意Load返回两个类型，分别是interface{}, bool
	// fmt.Println(v["province"]) // 注意 interface {} does not support indexing，所以这里报错
	// 解决方法：加一个类型断言
	fmt.Println(v.(map[string]string)["province"]) // 就可以正确输出 江苏
}
