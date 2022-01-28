package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	// select + channel
	func() {
		var wg sync.WaitGroup

		wg.Add(1)
		stopCh := make(chan bool)

		go func() {
			defer wg.Done()
			watchDog(stopCh, "|monitor dog|")
		}()
		time.Sleep(5 * time.Second)
		stopCh <- true
		wg.Wait()
	}()

	// select + context
	func() {
		var wg sync.WaitGroup
		wg.Add(1)
		ctx, stop := context.WithCancel(context.Background())
		go func() {
			defer wg.Done()
			watchDog2(ctx, "|Monitor dog two | ")
		}()

		time.Sleep(5 * time.Second)
		stop() // 发指令
		wg.Wait()
	}()

	// 1.Context 以上是 select + channel 常用使用场景
	//Context 就是用来简化解决这些问题的，并且是并发安全的。Context 是一个接口，它具备手动、定时、超时发出取消信号、传值等功能，
	//主要用于控制多个协程之间的协作，尤其是取消操作。一旦取消指令下达，那么被 Context 跟踪的这些协程都会收到取消信号，就可以做清理和退出操作。

	//Deadline 方法可以获取设置的截止时间，第一个返回值 deadline 是截止时间，到了这个时间点，Context 会自动发起取消请求，第二个返回值 ok 代表是否设置了截止时间。
	//Done 方法返回一个只读的 channel，类型为 struct{}。在协程中，如果该方法返回的 chan 可以读取，则意味着 Context 已经发起了取消信号。通过 Done 方法收到这个信号后，就可以做清理操作，然后退出协程，释放资源。
	//Err 方法返回取消的错误原因，即因为什么原因 Context 被取消。
	//Value 方法获取该 Context 上绑定的值，是一个键值对，所以要通过一个 key 才可以获取对应的值。

	// 我们不需要自己实现context, go 已经封装好了context
	//空 Context：不可取消，没有截止时间，主要用于 Context 树的根节点。
	//可取消的 Context：用于发出取消信号，当取消的时候，它的子 Context 也会取消。
	//可定时取消的 Context：多了一个定时的功能。
	//值 Context：用于存储一个 key-value 键值对。
}

func watchDog2(ctx context.Context, s string) {
	for true {
		select {
		case <-ctx.Done():
			fmt.Println(s, "Stop command has receive, Right now stop")
			return
		default:
			fmt.Println(s, "is Running")
		}
		time.Sleep(time.Second)
	}
}

func watchDog(ch chan bool, s string) {
	for true {
		select {
		case <-ch:
			fmt.Println(s, "must be to stop")
			return
		default:
			fmt.Println(s, "Monitor is running")
		}
		time.Sleep(time.Second)
	}
}
