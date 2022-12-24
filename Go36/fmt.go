/*
2022-12-23
fmt包常用函数：
	将内容输出到系统的标准输出：fmt.Print,fmt.Println
占位符	说明
%v	值的默认格式表示
%+v	类似%v，但输出结构体时会添加字段名
%#v	值的Go语法表示
%T	打印值的类型
%%	百分号

%w 错误
%t 布尔值
%p	指针，表示为十六进制，并加上前导的0x
整数：
%b	表示为二进制
%c	该值对应的unicode码值
%d	表示为十进制
%o	表示为八进制
%x	表示为十六进制，使用a-f
%X	表示为十六进制，使用A-F
%U	表示为Unicode格式：U+1234，等价于”U+%04X”
%q	该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
浮点数：
%b	无小数部分、二进制指数的科学计数法，如-123456p-78
%e	科学计数法，如-1234.456e+78
%E	科学计数法，如-1234.456E+78
%f	有小数部分但无指数部分，如123.456。%f，默认宽度，默认精度；%9f，宽度9，默认精度；%.2f默认宽度，精度2；%9.2f	宽度9，精度2；%9.f宽度9，精度0。
%F	等价于%f
%g	根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
%G	根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）
字符串：
%s	直接输出字符串或者[]byte
%q	该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
%x	每个字节用两字符十六进制数表示（使用a-f
%X	每个字节用两字符十六进制数表示（使用A-F）
*/
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func output1() { // 输出到标准输出
	fmt.Print("hello\n")               // 直接输出
	fmt.Printf("hello, %s\n", "world") // 格式化输出
	fmt.Println("hello")               // 自带换行，不支持格式化输出
}

func output2() { // 向指定文件输出
	// 比如指定文件为标准输出os.Stdout
	fmt.Fprintln(os.Stdout, "hello, %s", "world") // 除了F以外可以指定输出文件，其他特性一样
	// 打开一个文件，然后指定其为目标。这个涉及到os.OpenFile函数，稍后在os包测试文件中写。
}

func output3() {
	// 将传入的数据生成，返回一个字符串
	s1 := fmt.Sprint("hello string ")
	fmt.Print(s1)

	// 根据传入的参数数据生成格式化字符，并返回一个改字符串的错误
	e1 := fmt.Errorf("hello error ")
	fmt.Print(e1)
	e2 := errors.New("原始错误e")
	w := fmt.Errorf("Wrap了一个错误%w", e2)
	fmt.Println(w)
}

// 从标准输入获取用户输入
func input1() {
	var (
		name    string
		age     int
		married bool
	)
	fmt.Scan(&name, &age, &married) //将空白符(空格、回车符、换行符、水平制表符等)分割的数据分别存在指定的参数
	// Scanf从标准输入扫描文本,根据format参数指定的格式去读取由空白符分隔的值保存到传递本函数的参数中
	// fmt.Scanf("1:%v 2:%v 3:%v", &name, &age, &married)
	// Scanln类似Scan，它在遇到换行时才停止扫描。最后一个数据后面必须有换行或者到达结束位置。
	// fmt.Scanln(&name,&age,&married)
	fmt.Printf("扫描结果 name:%v age:%d marridL:%v", name, age, married)
}

func input2() {
	// 有时想获取完整输入内容,输入可能包含空格等。这种情况下可以使用bufio包来实现.
	reader := bufio.NewReader(os.Stdin) // 标准输出生成对象
	fmt.Print("请输入两行内容:")

	// ReadString从输入中读取数据，直到分隔符出现，返回包括分割符在内的字符串。
	// 如果ReadString在找到一个分割符之前就遇到了错误，它会返回已读取的数据和错误本身。
	// ReadString返回的 err!=nil，当且仅当它返回的数据不以分割符结尾。Scanner更加易于使用。
	text, _ := reader.ReadString('\n') // 读取到回车换行结束
	text2, _ := reader.ReadString('\n')

	// fmt.Printf("%v \n", text)      //打印输出
	fmt.Printf("%s 测试\n", text)    // 注意字符串里包含换行符，所以格式化字符串里可以不用多加'\n'再换行。很有意思的是%s是直接输出字符串，而字符串中包括回车符，所以%s后面的一个空格及后面内容会在下一行输出！
	text = strings.TrimSpace(text) // TrimSpace:去头部和尾部的空格、换行符。所以下面为了输出有换行，需要添加‘\n’
	// fmt.Printf("%v \n", text)      //打印输出
	fmt.Printf("%v \n", text2)

	// reader还有其他的读取方式。提醒：要掌握reader和writer的常用方法。
	text3, _, _ := reader.ReadLine() // 注意：方法不需要输入参数，返回参数有三个([]byte数组)！ReadLine是底层的行读取原语，不建议使用。
	//是一个低级的用于读取一行数据的方法，大多数调用者应该使用 ReadBytes('\n') 或者 ReadString('\n')。
	// ReadLine 返回一行，不包括结尾的回车字符，如果一行太长（超过缓冲区长度），参数 isPrefix 会设置为 true 并且只返回前面的数据，剩余的数据会在以后的调用中返回。
	fmt.Printf("%v \n", string(text3))
	text4, _, _ := reader.ReadLine()
	fmt.Printf("%v \n", string(text4))
}

func input3() {
	/* Scanner读取数据，输入。
	   Scanner提供了简单的读取数据的接口。连续地调用Scan方法将会跨到文件中的一个个特定符号，
	   并跳过特定符号之间的字符。这些符号由类型为SplitFunc的分割函数定义。默认的分割函数是
	   将输入按行切割成一块块，每一块结果不包括行结束符。分割函数定义在这个包的内部。用户可
	   以自定义分割函数。

		终止条件：linux：ctrl+D Windows：ctrl+z 回车
	*/
	c := make([]string, 0)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		s := input.Text()
		fmt.Printf("cur line is %s\n", s)
		c = append(c, s)
	}

	for n, line := range c {
		if n > 1 {
			fmt.Printf("%d  %s\n", n, line)
		}
	}
}

func main() {
	// output1()
	// output2()
	// output3()
	// input1()
	// input2()
	input3()
}
