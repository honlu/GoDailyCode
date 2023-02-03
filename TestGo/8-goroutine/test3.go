/*
更复杂的场景如何做并发控制呢？比如子协程中开启了新的子协程，或者需要同时控制多个子协程。这种场景下，select+chan的方式就显得力不从心了
Go 语言提供了 Context 标准库可以解决这类场景的问题，Context 的作用和它的名字很像，上下文，即子协程的下上文。
Context 有两个主要的功能：
	通知子协程退出（正常退出，超时退出等）；
	传递必要的参数。

大杀器：context包
Context通常被译作上下文，是一个比较抽象的概念。
在讨论链式调用技术时也经常会提到上下文。
一般理解为程序单元的一个运行状态、现场、快照，而翻译中上下又很好地诠释了其本质，上下则是存在上下层的传递，上会把内容传递给下。
在Go语言中，程序单元指的是Goroutine。

Context的创建和调用关系是层层递进的，也就是我们通常所说的链式调用，类似数据结构里的树.
从根节点开始，每一次调用就衍生一个叶子节点。
首先，生成根节点，使用context.Background方法生成.
而后可以进行链式调用使用context包里的各类方法，context包里的所有方法：
- func Background() Context  // 创建根 Context，通常在 main 函数、初始化和测试代码中创建，作为顶层 Context。
- func TODO() Context
- func WithCancel(parent Context) (ctx Context, cancel CancelFunc) // 创建ctx可取消的子 Context，同时返回函数，即可以主动通知子协程退出。
- func WithValue(parent Context, key, val interface{}) Context // context.WithValue() 创建了一个基于 ctx 的子 Context，并携带了值 options。在子协程中，使用 ctx.Value("options") 获取到传递的值，读取/修改该值。
- func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) // 创建具有超时退出机制的 Context 对象。
- func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc) // 超时退出可以控制子协程的最长执行时间，那 context.WithDeadline() 则可以控制子协程的最迟退出时间。

这里仅以WithCancel和WithValue方法为例来实现控制并发和通信。

有兴趣可以读一下context包的源码，会发现Context的实现其实是结合了Mutex锁和channel而实现的。
其实并发、同步的很多高级组件万变不离其宗，都是通过最底层的数据结构组装起来的，只要知晓了最基础的概念，上游的架构也可以一目了然。
*/

package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type favContextKey string

func main() {
	wg := &sync.WaitGroup{}
	values := []string{"https://www.baidu.com/", "https://www.zhihu.com/"}
	ctx, cancel := context.WithCancel(context.Background())

	for _, url := range values {
		wg.Add(1)
		subCtx := context.WithValue(ctx, favContextKey("url"), url)
		go reqURL(subCtx, wg)
	}

	go func() {
		time.Sleep(time.Second * 3)
		cancel()
	}()

	wg.Wait()
	fmt.Println("exit main goroutine")
}

func reqURL(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	url, _ := ctx.Value(favContextKey("url")).(string)
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("stop getting url:%s\n", url)
			return
		default:
			r, err := http.Get(url)
			if r.StatusCode == http.StatusOK && err == nil {
				body, _ := ioutil.ReadAll(r.Body)
				subCtx := context.WithValue(ctx, favContextKey("resp"), fmt.Sprintf("%s%x", url, md5.Sum(body)))
				wg.Add(1)
				go showResp(subCtx, wg)
			}
			r.Body.Close()
			//启动子goroutine是为了不阻塞当前goroutine，这里在实际场景中可以去执行其他逻辑，这里为了方便直接sleep一秒
			// doSometing()
			time.Sleep(time.Second * 1)
		}
	}
}

func showResp(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop showing resp")
			return
		default:
			//子goroutine里一般会处理一些IO任务，如读写数据库或者rpc调用，这里为了方便直接把数据打印
			fmt.Println("printing ", ctx.Value(favContextKey("resp")))
			time.Sleep(time.Second * 1)
		}
	}
}
