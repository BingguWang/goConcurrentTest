package main

import "fmt"

/*
	close函数关闭通道A来告知从该通道A接收值的goroutine停止等待。
	当通道被关闭时，往该通道发送值会引发panic，从该通道里接收的值一直都是类型零值。
*/
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1) //关闭的时候会通知下面的协程不要再等待了
	}()

	go func() {
		for {
			a, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- a * a
		}
		close(ch2)
	}()

	for v := range ch2 {
		fmt.Println(v)
	}
}
