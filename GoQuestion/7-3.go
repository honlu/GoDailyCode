package main

/*
2022-9-22
question:
	Go方法、变量、属性开头大小写的问题
answer：
	小写开头的方法、属性或 struct 是私有的，同样在json 解码或转码的时候无法对私有属性的转换。
	本例中是无法正常得到People的name值的。而且，私有属性name也不应该加json的标签。

	Go中 ``语法：不转译字符串的意思，可以json\yaml等一起使用
*/

import (
	"encoding/json"
	"fmt"
)

type People struct {
	name string `json:"name"` // 属性name,可以改为Name，本次输出结果不一样
}

func main() {
	js := `{
		"name":"11"
	}`
	var p People
	err := json.Unmarshal([]byte(js), &p) // 解码： json格式数据转换给Go类型的数据
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("people: ", p) // 如果Peope属性为name，则不进行转译，输出{}；若为Name，则输出{11}
}
