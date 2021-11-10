package main

import (
	"fmt"
	"time"
)

//  select 的多路复用，用于处理异步IO问题
func main() {
	//如果要从多个通道接收数据，可能会这样写
	// for {
	// 	data,ok:=<-ch1
	// 	data,ok:=<-ch2
	// }//但是在前面的语句接收到数据之前，前面的语句会阻塞到后面的语句，性能很差，select可以改良，同时响应多个通道操作！！

	/*
		select的每个case会对应一个接收或发送操作。
		select会一直等待，直到某个case的通信操作完成时，就会执行case分支对应的语句。完成select代码块的执行
		通过有多个case可以执行，会随机选一个case执行
	*/

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		ch1 <- "c1"
	}()
	go func() {
		time.Sleep(time.Second * 5)
		// time.Sleep(time.Second * 1)//如果改成1秒，就会看到随机选择case执行了
		ch1 <- "c2"
	}()

	select {
	case a1 := <-ch1:
		fmt.Println(a1)
	case a2 := <-ch2:
		fmt.Println(a2)
	} //因为c1总是会先完成case的操作，所以总是执行的是第一个case

}
