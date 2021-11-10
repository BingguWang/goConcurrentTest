package main

import (
	"fmt"
	"time"
)

//  select 的多路复用，用于处理异步IO问题
// select可以用用于判断管道是否存满
func main() {

	ch := make(chan int, 5)

	go func() { //创建一个写协程，让写的速度快过读取的速度导致通道满
		i := 1
		for {
			select {
			case ch <- i:
				fmt.Printf("通道还没满，写入数据%v\n", i)
			default:
				fmt.Println("通道满了,无法写入")
			}
			i++
			time.Sleep(time.Second * 1)
		}
	}()

	//取数据
	for da := range ch {
		fmt.Println("取出数据", da)
		time.Sleep(time.Second * 2)
	}

}
