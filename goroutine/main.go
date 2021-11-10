package main

import "fmt"

func SayHello() {
	fmt.Println("hello")
}

func main() {
	go SayHello()
	fmt.Println("main") //运行结果都是main，不会输出和hello
}

/*
	因为main函数默认会被新建一个goroutine,所以会最先启动的协程会是main函数
	SayHello建立协程需要时间，而main协程运行完后，其他的协程都会终止
*/
