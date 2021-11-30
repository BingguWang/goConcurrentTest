package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		time.Sleep(1 * time.Second)
		ch <- 1
		fmt.Println("--关闭通道")
		close(ch)
	}()
	for {
		select {
		case <-ch: // 读到信息
			time.Sleep(500 * time.Millisecond)
			fmt.Println("信息已放入通道，通道关闭，读到信息")
		default: // 读不到信息
			time.Sleep(500 * time.Millisecond)
			fmt.Println("信息还没放入通道，通道还没关闭，读不到信息")
		}
	}
	/**
	可以看到，在for select中，当通道关闭后会一直处于case <- ch: 这种奇怪的现象，<- ch在通道ch关闭后竟然还会一直读一直读！！！！
	个人认为这种现象是因为无缓冲通道的原因，因为读取和接收的及时性

	解决办法是，_, ok := <- ch 根据OK，OK == false了也就是通道关闭了，前面加上
	if !ok {
		ch = nil
		return
	}
	*/
}
