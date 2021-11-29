package main

import "fmt"

//缓冲通道

func main() {
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2 //因为是带缓存的，所以可以发多个数据
	//ch <- 3 // 但是如果缓存中没有了还在调用<-ch就会出现死锁，也就是说可以有余，但不能透支

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
