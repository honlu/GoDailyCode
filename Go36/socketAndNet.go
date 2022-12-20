/*
46. 访问网络服务
socket，常被翻译为套接字，它应该算是网络编程世界中最为核心的知识之一了。
所谓 socket，是一种 IPC 方法。IPC 是 Inter-Process Communication 的缩写，可以被翻译为进程间通信。
多个进程之间相互通信的方法主要包括：
	系统信号（signal）、管道（pipe）、套接字 （socket）、文件锁（file lock）、消息队列（message queue）、信号灯（semaphore，有的地方也称之为信号量）等。
对于 socket，Go 语言与之相应的程序实体都在其标准库的net代码包中。

syscall.Socket函数接受的那三个参数。都是int类型，分别是想要创建的 socket 实例通信域、类型以及使用的协议。
Socket 的通信域主要有这样几个可选项：IPv4 域、IPv6 域和 Unix 域。
三种通信域分别可以由syscall代码包中的常量AF_INET、AF_INET6和AF_UNIX表示。
Socket 的类型一共有 4 种，分别是：SOCK_DGRAM、SOCK_STREAM、SOCK_SEQPACKET以及SOCK_RAW。
通常，只要明确指定了前两个参数的值，我们就无需再去确定第三个参数值了，一般把它置为0就可以了。这时，内核程序会自行选择最合适的协议。
比如，当前两个参数值分别为syscall.AF_INET和syscall.SOCK_DGRAM的时候，内核程序会选择 UDP 作为协议。
又比如，在前两个参数值分别为syscall.AF_INET6和syscall.SOCK_STREAM时，内核程序可能会选择 TCP 作为协议。

*/
package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"runtime"
	"syscall"
	"time"
)

func syscall1() {
	fd1, err := syscall.Socket(
		syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	if err != nil {
		fmt.Printf("socket error: %v\n", err)
		return
	}
	defer syscall.Close(fd1)
	fmt.Printf("The file descriptor of socket：%d\n", fd1)

	// 省略若干代码。
	// 如果真要完全使用syscall包中的程序实体建立网络连接的话，
	// 过程太过繁琐而且完全没有必要。
	// 所以，我在这里就不做展示了。
}

func net1() {
	network := "tcp"
	host := "google.cn"
	reqStrTpl := `HEAD / HTTP/1.1
		Accept: */*
		Accept-Encoding: gzip, deflate
		Connection: keep-alive
		Host: %s
		User-Agent: Dialer/%s
		`

	// 示例1。
	network1 := network + "4"
	address1 := host + ":80"
	fmt.Printf("Dial %q with network %q ...\n", address1, network1)
	conn1, err := net.Dial(network1, address1)
	if err != nil {
		fmt.Printf("dial error: %v\n", err)
		return
	}
	defer conn1.Close()

	reqStr1 := fmt.Sprintf(reqStrTpl, host, runtime.Version())
	fmt.Printf("The request:\n%s\n", reqStr1)
	_, err = io.WriteString(conn1, reqStr1)
	if err != nil {
		fmt.Printf("write error: %v\n", err)
		return
	}
	fmt.Println()

	reader1 := bufio.NewReader(conn1)
	line1, err := reader1.ReadString('\n')
	if err != nil {
		fmt.Printf("read error: %v\n", err)
		return
	}
	fmt.Printf("The first line of response:\n%s\n", line1)
	fmt.Println()

	// 示例2。
	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		MinVersion:         tls.VersionTLS10,
	}
	network2 := network
	address2 := host + ":443"
	fmt.Printf("Dial %q with network %q ...\n", address2, network2)
	conn2, err := tls.Dial(network2, address2, tlsConf)
	if err != nil {
		fmt.Printf("dial error: %v\n", err)
		return
	}
	defer conn2.Close()

	reqStr2 := fmt.Sprintf(reqStrTpl, host, runtime.Version())
	fmt.Printf("The request:\n%s\n", reqStr2)
	_, err = io.WriteString(conn2, reqStr2)
	if err != nil {
		fmt.Printf("write error: %v\n", err)
		return
	}

	reader2 := bufio.NewReader(conn2)
	line2, err := reader2.ReadString('\n')
	if err != nil {
		fmt.Printf("read error: %v\n", err)
		return
	}
	fmt.Printf("The first line of response:\n%s\n", line2)
	fmt.Println()
}

type dailArgs struct {
	network string
	address string
	timeout time.Duration
}

func net2() {
	dialArgsList := []dailArgs{
		{
			"tcp",
			"google.cn:80",
			time.Millisecond * 500,
		},
		{
			"tcp",
			"google.com:80",
			time.Second * 2,
		},
		{
			// 如果在这种情况下发生的错误是：
			// "connect: operation timed out"，
			// 那么代表着什么呢？
			//
			// 简单来说，此错误表示底层的socket在连接网络服务的时候先超时了。
			// 这时抛出的其实是'syscall.ETIMEDOUT'常量代表的错误值。
			"tcp",
			"google.com:80",
			time.Minute * 4,
		},
	}
	for _, args := range dialArgsList {
		fmt.Printf("Dial %q with network %q and timeout %s ...\n",
			args.address, args.network, args.timeout)
		ts1 := time.Now()
		conn, err := net.DialTimeout(args.network, args.address, args.timeout) // 给定超时时间：代表着函数为网络连接建立完成而等待的最长时间。这是一个相对的时间。它会由这个函数的参数timeout的值表示。
		ts2 := time.Now()
		fmt.Printf("Elapsed time: %s\n", time.Duration(ts2.Sub(ts1)))
		if err != nil {
			fmt.Printf("dial error: %v\n", err)
			fmt.Println()
			continue
		}
		defer conn.Close()
		fmt.Printf("The local address: %s\n", conn.LocalAddr())
		fmt.Printf("The remote address: %s\n", conn.RemoteAddr())
		fmt.Println()
	}
}

func main() {
	net1()
}
