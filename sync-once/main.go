package main

import (
	"fmt"
	"sync"
	"time"
)

//sync.Once
/*
	用于解决一次性初始化问题，作用类似init函数，只执行一次

	和init函数的区别：
	init:	文件包首次被加载时执行，只执行一次
	sync.Once:	代码运行中真正用到的时候才执行，只执行一次

	高并发场景中，需要确保某些操作只执行一次，比如只加载一次配置文件，只关闭一次通道

	sync.Once.Do(f)操作是并发安全的！！！

	type Once struct {
		done uint32 //记录是否已完成过初始化，保证值进行一次初始化
		m    Mutex //互斥锁，保证初始化操作的并发安全
	}
	//Once只有一个方法Do，需要传入需要执行的方法
*/

func main() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("只会打印一次")
	}
	// done := make(chan bool)
	for i := 0; i < 6; i++ {
		go func() {
			once.Do(onceBody) //传入需要执行的方法
			// done <- true      //已执行过
		}()
	}

	time.Sleep(time.Second * 2)
	fmt.Printf("%#v\n", once) //可以看到执行了一次后once实例的done就变成了true了
	fmt.Printf("%s\n", once)  //可以看到执行了一次后once实例的done就变成了true了

	// for i := 0; i < 6; i++ {
	// 	<-done
	// }

	for {
	}
}
