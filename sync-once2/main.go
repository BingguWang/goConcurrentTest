package main

import (
	"fmt"
	"sync"
)

//sync.Once可以用于只关闭一次通道，防止对已关闭的通道关闭导致的宕机

var once sync.Once
var wait sync.WaitGroup

func main() {
	ch1 := make(chan int, 10) //双向通道可以转为单向通道
	ch2 := make(chan int, 10) //双向通道可以转为单向通道

	wait.Add(3)

	go f1(ch1)
	go f2(ch1, ch2)
	go f2(ch1, ch2)

	wait.Wait()
	for result := range ch2 {
		fmt.Println(result)
	}
}

func f1(ch chan<- int) { //写数据入通道
	defer wait.Done()
	for i := 0; i < 10; i++ {
		ch <- i
	}
	fmt.Println("关闭ch1")
	close(ch)
}

func f2(ch1 <-chan int, ch2 chan<- int) { //从通道取数据
	defer wait.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- 2 * x
	}
	once.Do(func() {
		fmt.Println("关闭ch2")
		close(ch2) //这样只会关一次通道ch2,其他协程不会再关
	})
}
