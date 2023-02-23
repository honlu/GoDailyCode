/*
go-zero社区别人发的代码
2023-2-2 by lu update

一下逻辑太有问题了，没改通。
*/
package main

const total = 10 // 1000000000

/*
类型chan<- int表示一个只发送int的channel，只能发送不能接收。
相反，类型<-chan int表示一个只接收int的channel，只能接收不能发送。
（箭头<-和关键字chan的相对位置表明了channel的方向。）这种限制将在编译期检测。
*/
