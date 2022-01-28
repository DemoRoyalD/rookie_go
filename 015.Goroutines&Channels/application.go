package main

import (
	"fmt"
	"time"
)

func main() {
	moreDefer()

	// 1. goroutines
	go moreDefer() // 启动线程，必须阻塞主main线程，不然会看不到 子线程

	// 2. channel
	channel := make(chan string)
	go func() {
		fmt.Println("Dongchen")
		channel <- "dongchen achieve the task"
	}()

	fmt.Println("=================")

	result := <-channel
	fmt.Println(" result : ", result)
	time.Sleep(time.Second)
	// 3. 上面创建的channel 是没有缓冲的
	cacheChannel := make(chan int, 5)
	go func() {
		fmt.Println("cache channel start")
		cacheChannel <- 20000000
		cacheChannel <- 30000000
		cacheChannel <- 40000000
		cacheChannel <- 50000000
		cacheChannel <- 60000000
		cacheChannel <- 70000000
	}()
	fmt.Println("Cache channel size ", len(cacheChannel))
	fmt.Println("Cache channel capacity ", cap(cacheChannel))
	fmt.Println(<-cacheChannel)
	fmt.Println(<-cacheChannel)
	fmt.Println(<-cacheChannel)
	fmt.Println(<-cacheChannel)
	fmt.Println(<-cacheChannel)
	cacheResult := <-cacheChannel
	fmt.Println(cacheResult)
	close(cacheChannel)
	close(channel)

	// 4.单向channel  var ch1 = make(chan<- int) :只接受 var ch2 = make(<-chan int) 只传出

}

func moreDefer() {
	defer fmt.Println("First defer")

	defer fmt.Println("Second defer")

	defer fmt.Println("Third defer")

	fmt.Println("Function self coe")
}
