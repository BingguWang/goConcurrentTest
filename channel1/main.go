package main

/*
>>>>>>>>>>>>>channel简单通信

*/
import "fmt"

func main() {

	ch := make(chan string)
	go func() {
		fmt.Println("新协程开始运行")
		ch <- "我结束运行了"
		fmt.Println("新协程结束运行")
	}()
	fmt.Println("main在等新协程运行完毕")
	<-ch //会在收到新协程写入通道的消息之前一直阻塞
	fmt.Println("main运行完毕")

}
