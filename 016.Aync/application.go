package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	sum   int
	mutex sync.Mutex
)

func add(i int) {
	sum += i
}

func add2(i int) {
	mutex.Lock()
	sum += i
	mutex.Unlock()
}

func main() {

	//for i := 0; i < 100; i++ {
	//	go add(10)
	//}
	time.Sleep(2 * time.Second)
	fmt.Println("sum is ", sum)

	// 1. sync.Mutex
	for i := 0; i < 100; i++ {
		go add2(10)
	}
	//var tmp = 0
	//for i := 0; i < 100; i++ {
	//	tmp += 10
	//}
	//fmt.Println("tmp is ", tmp)
	fmt.Println("sum is ", sum)
	time.Sleep(10 * time.Second)

	// 2. sync.WaitGroup   因为异步的话主线程退出就看不到子线程的工作情况，所以有了 sync.WaitGroup
	var wg sync.WaitGroup
	wg.Add(110) // 计数器
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()                   // 计数器减一  defer: 类似c++中那个啥来着，解析函数？ no ,忘记了
			fmt.Println("Hello index -> ", i) // 这里数据会有问题，因为i 一直在变，但是输出走的时候并发，所以数据并不是想象之中的样子
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			fmt.Println("World index -> ", i)
		}()
	}

	wg.Wait() // 一直等待，只要计数器为0

	// 3. sync.Once  让代码只执行一次，哪怕是在高并发的情况下

	doOnce()
}

func doOnce() {
	var once sync.Once

	onceBody := func() {
		fmt.Println(" Only once")
	}
	done := make(chan bool)

	// 启动10个协诚执行once.Do(onceBody)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}
