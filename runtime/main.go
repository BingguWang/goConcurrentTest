package main

import (
	"fmt"
	"runtime"
)

/*
	>>>>>>>>>>>>>>>runtime包
*/
// func main() {
// 	go func(s string) {
// 		for i := 0; i < 2; i++ {
// 			fmt.Println(s)
// 		}
// 	}("wb")
// 	//runtime.Gosched()//让出CPU时间片，重新等待安排任务
// 	for i := 0; i < 2; i++ {
// 		runtime.Gosched() //有点类似java中的jork
// 		fmt.Println("main")
// 	}
// }

func main() {

	go func() {
		defer fmt.Println("a.defer")
		func() {
			defer fmt.Println("b.defer")
			defer fmt.Println("bb.defer")
			//结束协程
			runtime.Goexit() //结束了不会执行下面的内容，但是会执行延迟函数的内容,当然也只是执行此语句前的defer函数
			defer fmt.Println("c.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()

	for {

	}
}

//另外，    runtime.GOMAXPROCS(1)可以设置CPU核数
