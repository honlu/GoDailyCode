/*
2022-9-22
question:
	有一个机器人，给一串指令，L左转，R右转，F前进一步，B后退一步，问最后机器人的坐标，最开始，机器人位于 0 0，方向为正Y。
	可以输入重复指令n ： 比如 R2(LF) 这个等于指令 RLFLF。
	问最后机器人的坐标是多少？
answer:
	这里的一个难点是解析重复指令和移动坐标公式。指令解析成功，计算坐标就简单了。
	想象一个x,y坐标轴，机器人在远点，方向y轴正方向。

	x表示x坐标，y表示y坐标，z表示当前方向。 L、R 命令会改变值z，F、B命令会改变值x、y。 值x、y的改变还受当前的z值影响。
	如果是重复指令，那么将重复次数和重复的指令存起来递归调用即可。
*/
package main

import (
	"fmt"
	"unicode"
)

const (
	Left  = iota // iota是从0开始，每一行都是往下递增。 左 0
	Up           // 上 1
	Right        // 右 2
	Down         // 下 3
)

func run(s string, x0, y0, z0 int) (x, y, z int) {
	x, y, z = x0, y0, z0
	repeat := 0
	repeatS := ""

	for _, val := range s {
		switch {
		case unicode.IsNumber(val): // 数字可能是几位数
			repeat = repeat*10 + int(val-'0')
		case repeat > 0 && val != '(' && val != ')': // 当数字出现后，要保存重复的字母
			repeatS = repeatS + string(val)
		case val == ')':
			for i := 0; i < repeat; i++ {
				x, y, z = run(repeatS, x, y, z)
			}
			repeat = 0
			repeatS = ""
		case val == 'L': // 方向更改: 左转
			z = (z - 1 + 4) % 4
		case val == 'R': // 右转
			z = (z + 1) % 4
		case val == 'F': // 向前一步
			switch {
			case z == Left || z == Right:
				x = x + z - 1
			case z == Up || z == Down: // Up:1 Down:3
				y = y - z + 2
			}

		case val == 'B': // 后退一步
			switch {
			case z == Left || z == Right: // Left:0, Right:2
				x = x - z + 1
			case z == Up || z == Down: // Up:1 Down:3
				y = y + z - 2
			}
		}

	}
	return x, y, z
}

func main() {
	fmt.Println(run("R2(LF)", 0, 0, Up)) // LFLF
}
